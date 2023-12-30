package main

import (
	"tgotify/api"
	tgClient "tgotify/client"
	"tgotify/config"
	"tgotify/consumer"
	"tgotify/processors/telegram"
	"tgotify/storage/postgres"

	"github.com/sirupsen/logrus"
)

func main() {
	// Initialize the configuration from the config package.
	cfg := config.Init()

	// Create a new PostgreSQL database connection using the configuration.
	db := postgres.New(&cfg.Storage)

	tokens, err := db.EnabledTokens()
	if err != nil {
		logrus.Fatal("Unable to get tokens from db", err)
	}
	tg := tgClient.New(cfg.TgHost, tokens, db)

	processor := telegram.New(tg, db)

	const batchSize = 100
	consumer := consumer.New(processor, processor, batchSize)

	go func() {
		consumer.Start()
	}()

	// Create and start the API server.
	api.CreateRouter(tg, db)
}
