package config

type Config struct {
	EnabledRules   []string `json:"enabled-rules,omitempty"`
	SensitiveWords []string `json:"sensitive-words,omitempty"`
	AllowedSymbols string   `json:"allowed-symbols,omitempty"`
}
