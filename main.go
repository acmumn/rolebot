package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	token := os.Getenv("BOT_TOKEN")

	bot := newBot(token)
	bot.Run()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	log.Println("shutting down")
	bot.Close()
}
