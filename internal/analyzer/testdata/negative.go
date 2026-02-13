package testdata

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample()
	sugar := logger.Sugar()

	slog.Debug("Debug message") // want "log message should not be capitalized"
	slog.Info("Info message")   // want "log message should not be capitalized"
	slog.Warn("Warn message")   // want "log message should not be capitalized"
	slog.Error("Error message") // want "log message should not be capitalized"

	logger.Debug(`Debug message`)   // want "log message should not be capitalized"
	logger.Info(`Info message`)     // want "log message should not be capitalized"
	logger.Warn("Warn message")     // want "log message should not be capitalized"
	logger.Error("Error message")   // want "log message should not be capitalized"
	logger.DPanic("DPanic message") // want "log message should not be capitalized"
	logger.Panic("Panic message")   // want "log message should not be capitalized"
	logger.Fatal("Fatal message")   // want "log message should not be capitalized"

	sugar.Debug("Debug message")           // want "log message should not be capitalized"
	sugar.Debugf("Debugf %s", "message")   // want "log message should not be capitalized"
	sugar.Debugw("Debugw message")         // want "log message should not be capitalized"
	sugar.Debugln("Debugln message")       // want "log message should not be capitalized"
	sugar.Info("Info message")             // want "log message should not be capitalized"
	sugar.Infof("Infof %s", "message")     // want "log message should not be capitalized"
	sugar.Infow("Infow message")           // want "log message should not be capitalized"
	sugar.Infoln("Infoln message")         // want "log message should not be capitalized"
	sugar.Warn("Warn message")             // want "log message should not be capitalized"
	sugar.Warnf("Warnf %s", "message")     // want "log message should not be capitalized"
	sugar.Warnw("Warnw message")           // want "log message should not be capitalized"
	sugar.Warnln("Warnln message")         // want "log message should not be capitalized"
	sugar.Error("Error message")           // want "log message should not be capitalized"
	sugar.Errorf("Errorf %s", "message")   // want "log message should not be capitalized"
	sugar.Errorw("Errorw message")         // want "log message should not be capitalized"
	sugar.Errorln("Errorln message")       // want "log message should not be capitalized"
	sugar.DPanic("DPanic message")         // want "log message should not be capitalized"
	sugar.DPanicf("DPanicf %s", "message") // want "log message should not be capitalized"
	sugar.DPanicw("DPanicw message")       // want "log message should not be capitalized"
	sugar.DPanicln("DPanicln message")     // want "log message should not be capitalized"
	sugar.Panic("Panic message")           // want "log message should not be capitalized"
	sugar.Panicf("Panicf %s", "message")   // want "log message should not be capitalized"
	sugar.Panicw("Panicw message")         // want "log message should not be capitalized"
	sugar.Panicln("Panicln message")       // want "log message should not be capitalized"
	sugar.Fatal("Fatal message")           // want "log message should not be capitalized"
	sugar.Fatalf("Fatalf %s", "message")   // want "log message should not be capitalized"
	sugar.Fatalw("Fatalw message")         // want "log message should not be capitalized"
	sugar.Fatalln("Fatalln message")       // want "log message should not be capitalized"

	zap.L().Info("Info message")         // want "log message should not be capitalized"
	zap.S().Infof("Infof %s", "message") // want "log message should not be capitalized"

	slog.Info("русское сообщение")  // want "log message should be in english"
	logger.Info("кириллица")        // want "log message should be in english"
	sugar.Infof("привет %s", "мир") // want "log message should be in english"

	slog.Info("server-error")  // want "log message should contains only letters, numbers and spaces"
	logger.Info("user@email")  // want "log message should contains only letters, numbers and spaces"
	sugar.Info("start🚀server") // want "log message should contains only letters, numbers and spaces"

	token := "abc"
	bearer := "jwt"
	apiKey := "12345"
	authToken := "xxx"

	logger.Info("request",
		zap.String("key", token), // want "log message should not contains sensitive values"
	)

	slog.Debug(
		"debug message",
		slog.Int64("key", 5),
		slog.Any("key", apiKey), // want "log message should not contains sensitive values"
	)

	sugar.Infow("login",
		"cred", bearer, // want "log message should not contains sensitive values"
	)

	slog.Info("user data",
		"field", apiKey, // want "log message should not contains sensitive values"
	)

	logger.Info("auth",
		zap.String("user", authToken), // want "log message should not contains sensitive values"
	)
}
