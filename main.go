package main

import (
	"flag"

	"github.com/speedy/bot"
	"github.com/speedy/config"
	"github.com/speedy/db"
	"github.com/speedy/job"
)

func main() {
	var configFilePath string

	flag.StringVar(&configFilePath, "config", "config.toml", "Configuration file path")
	flag.Parse()

	config.Init(configFilePath)
	config.InitLogging()

	db.Init()
	db.Migrate(db.DB)
	defer db.Close()

	go job.Jobs()

	bot.Start(config.Env.Telegram.Token)
}
