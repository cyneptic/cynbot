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

	t := "MTE1MzMxNTUzODUyOTk2NDAzMg.GOVtZM.gyAobrDBt3ZKxQOsyABfRzwM0PfhPtkPqngxrk"
	gid := "1042387296613322763"

	b := bot.CreateNewBot(t, gid)
	defer b.S.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
