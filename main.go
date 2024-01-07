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
	cfg := config.Init()

	db := postgres.New(&cfg.Storage)

	tokens, err := db.EnabledClients()
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

	api.CreateRouter(tg, db)
}
