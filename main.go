package main

import (
	"flag"

	"github.com/SpeedTest-VF/bot"
	"github.com/SpeedTest-VF/config"
	"github.com/SpeedTest-VF/db"
	"github.com/SpeedTest-VF/job"
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
