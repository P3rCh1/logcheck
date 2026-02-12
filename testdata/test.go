package main

import (
	"log/slog"
)

func main() {
	log := slog.New(slog.DiscardHandler)
	log.Info("Starting server")
	slog.Info("Starting server")
}
