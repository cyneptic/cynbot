package main

import (
	"os"
	"os/signal"

	"github.com/cyneptic/cynbot/bot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	t := os.Getenv("BOT_TOKEN")
	gid := os.Getenv("GUILD_ID")

	b := bot.CreateNewBot(t, gid)
	defer b.S.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
