package server

import (
	"context"
	"log"
	conf "marilyn_manson_bot/config"
	"marilyn_manson_bot/pkg/logger"
	"marilyn_manson_bot/pkg/postgres"
	"time"

	repo "marilyn_manson_bot/internal/repository"

	tele "gopkg.in/telebot.v4"
)

var (
	Bot        *tele.Bot
	Log        logger.Logger
	Repository repo.DebtRepository
)

func init() {
	config, err := conf.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	Log, err = logger.NewLogrusLogger(config.LogLevel)
	if err != nil {
		log.Fatal("venom")
	}
	pg, err := postgres.New(context.Background(), config.PgConnectURL, postgres.WithLogger(Log))
	if err != nil {
		log.Fatal("Cant connect pg")
	}
	Repository = repo.NewDebtRepo(pg, Log)
	pref := tele.Settings{
		Token:  config.Token,
		Poller: &tele.LongPoller{Timeout: time.Duration(config.PoolerTimeout) * time.Second},
	}
	Bot, err = tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	// add handlers here
	Bot.Handle("/hello", helloHandler)
	Bot.Handle("/list", listOfDebts)
	Bot.Handle("/add", createDebt)
}
