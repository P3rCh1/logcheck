package main

import (
	"log/slog"
)

func tokEn123() int {
	return 0
}

type User struct {
	api_keyS int
}

func main() {

	log := slog.New(slog.DiscardHandler)
	var secret string
	log.Info("Starting server"+secret, tokEn123())
	log.Info("starting server"+"token", tokEn123())
	slog.Info("Starting server", "", func() int { return 0 }())
	slog.Info("starting server", "", User{}.api_keyS)
}
