package bot

import (
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
)

func Start(token string) {
	b, err := telebot.SpeedTest(telebot.Settings{
		Token:  2037914959:AAGKtCIGkTC6Whwj0wEZROPui86Tgk3_vsU,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	h := Handler{Bot: b}
	b.Handle("/start", h.Start)
	b.Handle("/sunucular", h.Servers)
	b.Handle("/test", h.Test)
	b.Handle("/yardim", h.Help)
	b.Handle("/son", h.Last)
	b.Handle("/eniyi", h.Best)
	b.Handle("/enkotu", h.Worst)
	b.Handle(telebot.OnText, h.Text)

	logrus.Info("Telegram bot başlıyor")
	b.Start()
}
