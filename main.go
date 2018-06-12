package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/thehowl/conf"
)

type Config struct {
	Token string `description:"Bot token for authentication"`
}

var defaultCfg = Config{}

func main() {
	configFile := flag.String("conf", "rolebot.conf", "config file location")
	flag.Parse()

	config := Config{}
	err := conf.Load(&config, *configFile)
	if err == conf.ErrNoFile {
		conf.Export(defaultCfg, *configFile)
		fmt.Println("Default configuration written to " + *configFile)
		os.Exit(0)
	}

	bot := newBot(config)
	bot.Run()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	log.Println("shutting down")
	bot.Close()
}
