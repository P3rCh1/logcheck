package rules

import (
	"go/token"
	"testing"

	"github.com/P3rCh1/logcheck/internal/utils"
	"github.com/cloudflare/ahocorasick"
	"github.com/stretchr/testify/assert"
)

type testcase struct {
	name    string
	logInfo *utils.LogInfo
	expMsg  string
	expPos  token.Pos
}

func TestCheckLowercase(t *testing.T) {
	tests := []testcase{
		{
			name: "empty message parts",
			logInfo: &utils.LogInfo{
				MsgParts: []utils.ItemAST{},
			},
		},
		{
			name: "message starts with lowercase",
			logInfo: &utils.LogInfo{
				MsgParts: []utils.ItemAST{
					{
						Data: "starting server",
						Pos:  token.Pos(100),
					},
				},
			},
		},
		{
			name: "message starts with uppercase",
			logInfo: &utils.LogInfo{
				MsgParts: []utils.ItemAST{
					{
						Data: "Starting server",
						Pos:  token.Pos(200),
					},
				},
			},
			expMsg: CapitalizedReport,
			expPos: token.Pos(200),
		},
		{
			name: "first part empty, second starts with uppercase",
			logInfo: &utils.LogInfo{
				MsgParts: []utils.ItemAST{
					{
						Data: "",
						Pos:  token.Pos(100),
					},
					{
						Data: "Failed to connect",
						Pos:  token.Pos(300),
					},
				},
			},
			expMsg: CapitalizedReport,
			expPos: token.Pos(300),
		},
		{
			name: "first part empty, second starts with lowercase",
			logInfo: &utils.LogInfo{
				MsgParts: []utils.ItemAST{
					{
						Data: "", Pos: token.Pos(100),
					},
					{
						Data: "failed to connect", Pos: token.Pos(300),
					},
				},
			},
		},
		{
			name: "starts with non-letter (digit)",
			logInfo: &utils.LogInfo{
				MsgParts: []utils.ItemAST{
					{
						Data: "123 starting server",
						Pos:  token.Pos(400),
					},
				},
			},
		},
		{
			name: "starts with non-letter (symbol)",
			logInfo: &utils.LogInfo{
				MsgParts: []utils.ItemAST{
					{
						Data: "[INFO] starting server",
						Pos:  token.Pos(500),
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			diag := CheckLowercase(test.logInfo)

			if test.expMsg == "" {
				assert.Nil(t, diag, "expected no diagnostic")
			} else {
				assert.NotNil(t, diag, "expected diagnostic")
				if diag != nil {
					assert.Equal(t, test.expMsg, diag.Message)
					assert.Equal(t, test.expPos, diag.Pos)
				}
			}
		})
	}
}

func TestCheckEnglish(t *testing.T) {
	tests := []testcase{
		{
			name: "empty message parts",
			logInfo: &utils.LogInfo{
				MsgParts: []utils.ItemAST{},
			},
		},
		{
			name: "english message",
			logInfo: &utils.LogInfo{
				MsgParts: []utils.ItemAST{
					{
						Data: "starting",
						Pos:  token.Pos(100),
					},
					{
						Data: "server",
						Pos:  token.Pos(200),
					},
				},
			},
		},
		{
			name: "with russian",
			logInfo: &utils.LogInfo{
				MsgParts: []utils.ItemAST{
					{
						Data: "starting",
						Pos:  token.Pos(100),
					},
					{
						Data: "сервер",
						Pos:  token.Pos(200),
					},
				},
			},
			expMsg: NotEnglishReport,
			expPos: token.Pos(200),
		},
		{
			name: "with over symbols",
			logInfo: &utils.LogInfo{
				MsgParts: []utils.ItemAST{
					{
						Data: "star-ting",
						Pos:  token.Pos(100),
					},
					{
						Data: "ser🙃ver",
						Pos:  token.Pos(200),
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			diag := CheckEnglish(test.logInfo)

			if test.expMsg == "" {
				assert.Nil(t, diag, "expected no diagnostic")
			} else {
				assert.NotNil(t, diag, "expected diagnostic")
				if diag != nil {
					assert.Equal(t, test.expMsg, diag.Message)
					assert.Equal(t, test.expPos, diag.Pos)
				}
			}
		})
	}
}

func TestCheckSensitiveLeak(t *testing.T) {
	tests := []testcase{
		{
			name: "empty message parts",
			logInfo: &utils.LogInfo{
				ArgNames: []utils.ItemAST{},
			},
		},
		{
			name: "without leaks",
			logInfo: &utils.LogInfo{
				ArgNames: []utils.ItemAST{
					{
						Data: "username",
						Pos:  token.Pos(100),
					},
					{
						Data: "err",
						Pos:  token.Pos(200),
					},
				},
			},
		},
		{
			name: "with leeks full",
			logInfo: &utils.LogInfo{
				ArgNames: []utils.ItemAST{
					{
						Data: "user",
						Pos:  token.Pos(100),
					},
					{
						Data: "token",
						Pos:  token.Pos(200),
					},
				},
			},
			expMsg: SensitiveLeakReport,
			expPos: token.Pos(200),
		},
		{
			name: "with leaks part",
			logInfo: &utils.LogInfo{
				ArgNames: []utils.ItemAST{
					{
						Data: "userPwdHash",
						Pos:  token.Pos(100),
					},
					{
						Data: "error",
						Pos:  token.Pos(200),
					},
				},
			},
			expMsg: SensitiveLeakReport,
			expPos: token.Pos(100),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			diag := CheckSensitiveLeak(
				ahocorasick.NewStringMatcher(SensitiveWordsDefault),
				test.logInfo,
			)

			if test.expMsg == "" {
				assert.Nil(t, diag, "expected no diagnostic")
			} else {
				assert.NotNil(t, diag, "expected diagnostic")
				if diag != nil {
					assert.Equal(t, test.expMsg, diag.Message)
					assert.Equal(t, test.expPos, diag.Pos)
				}
			}
		})
	}
}

func TestCheckSymbolsAndEmoji(t *testing.T) {
	tests := []testcase{
		{
			name: "empty message parts",
			logInfo: &utils.LogInfo{
				MsgParts: []utils.ItemAST{},
			},
		},
		{
			name: "english message",
			logInfo: &utils.LogInfo{
				MsgParts: []utils.ItemAST{
					{
						Data: "starting1",
						Pos:  token.Pos(100),
					},
					{
						Data: "2server",
						Pos:  token.Pos(200),
					},
				},
			},
		},
		{
			name: "with allowed",
			logInfo: &utils.LogInfo{
				MsgParts: []utils.ItemAST{
					{
						Data: "starting%%",
						Pos:  token.Pos(100),
					},
					{
						Data: "сервер 1.0",
						Pos:  token.Pos(200),
					},
				},
			},
		},
		{
			name: "with symbol",
			logInfo: &utils.LogInfo{
				MsgParts: []utils.ItemAST{
					{
						Data: "star-ting",
						Pos:  token.Pos(100),
					},
					{
						Data: "server",
						Pos:  token.Pos(200),
					},
				},
			},
			expMsg: SymbolOrEmojiReport,
			expPos: token.Pos(100),
		},
		{
			name: "with symbol and emoji",
			logInfo: &utils.LogInfo{
				MsgParts: []utils.ItemAST{
					{
						Data: "star🙃ting",
						Pos:  token.Pos(100),
					},
					{
						Data: "ser-ver",
						Pos:  token.Pos(200),
					},
				},
			},
			expMsg: SymbolOrEmojiReport,
			expPos: token.Pos(100),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			diag := CheckNoSymbolsAndEmoji(map[rune]struct{}{'%': {}, '.': {}}, test.logInfo)

			if test.expMsg == "" {
				assert.Nil(t, diag, "expected no diagnostic")
			} else {
				assert.NotNil(t, diag, "expected diagnostic")
				if diag != nil {
					assert.Equal(t, test.expMsg, diag.Message)
					assert.Equal(t, test.expPos, diag.Pos)
				}
			}
		})
	}
}
