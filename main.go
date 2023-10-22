package main

import (
	"tgotify/api"
	telegram "tgotify/client"
	"tgotify/config"
	"tgotify/storage/postgres"
)

func main() {
	// Initialize the configuration from the config package.
	cfg := config.Init()

	// Create a new PostgreSQL database connection using the configuration.
	db := postgres.New(&cfg.Storage)

	// Create a new Telegram client.
	tg := telegram.New(cfg.TgHost)

	// Create and start the API server.
	api.CreateRouter(tg, db)
}
