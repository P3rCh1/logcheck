package testdata

import (
	"log/slog"

	"go.uber.org/zap"
)

func Positive() {
	logger := zap.NewExample()
	sugar := logger.Sugar()
	slogLogger := slog.Default()

	slogLogger.Debug("debug message")
	slogLogger.Info("info message")
	slogLogger.Warn("warn message")
	slogLogger.Error("error message")

	slog.Debug("debug message")
	slog.Info("info message")
	slog.Warn("warn message")
	slog.Error("error message")

	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")
	logger.DPanic("dpanic message")
	logger.Panic("panic message")
	logger.Fatal("fatal message")

	sugar.Debug("debug message")
	sugar.Debugf("debug %s", "message")
	sugar.Debugw("debug message")
	sugar.Debugln("debug message")
	sugar.Info("info message")
	sugar.Infof("info %s", "message")
	sugar.Infow("info message")
	sugar.Infoln("info message")
	sugar.Warn("warn message")
	sugar.Warnf("warn %s", "message")
	sugar.Warnw("warn message")
	sugar.Warnln("warn message")
	sugar.Error("error message")
	sugar.Errorf("error %s", "message")
	sugar.Errorw("error message")
	sugar.Errorln("error message")
	sugar.DPanic("dpanic message")
	sugar.DPanicf("dpanic %s", "message")
	sugar.DPanicw("dpanic message")
	sugar.DPanicln("dpanic message")
	sugar.Panic("panic message")
	sugar.Panicf("panic %s", "message")
	sugar.Panicw("panic message")
	sugar.Panicln("panic message")
	sugar.Fatal("fatal message")
	sugar.Fatalf("fatal %s", "message")
	sugar.Fatalw("fatal message")
	sugar.Fatalln("fatal message")

	zap.L().Info("info message")
	zap.S().Infof("info %s", "message")

	slog.Info("attempt 3 of 5")
	logger.Info("processed 100 items")
	sugar.Info("ratio")

	logger.Info("user action",
		zap.String("username", "alice"),
		zap.String("action", "click"),
		zap.Int("code", 200),
	)

	sugar.Infow("request processed",
		"method", "GET",
		"status", 200,
		"duration", "1.5s",
	)

	slog.Info("operation completed",
		"operation", "sync",
		"success", true,
	)

	slog.Info("starting background worker pool with 10 workers")
	logger.Debug("cache hit ratio requests")
	sugar.Infof("user %s performed %d actions", "bob", 42)

	logger.Debug("debugging connection pool")
	logger.Info("server listening on 8080")
	logger.Warn("high memory usage")
	logger.Error("failed to connect to database")

	sugar.Debugw("query executed",
		"query", "SELECT * FROM users",
		"duration", "50ms",
		"rows", 5,
	)

	users := "users"

	logger.Info("map test",
		zap.Any(users, map[string]string{"alice": "active"}),
	)

	sugar.Infow("slice test", "user_ids", []int{1, 2, 3})
}
