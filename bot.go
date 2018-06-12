package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	dg *discordgo.Session
}

func newBot(config Config) (bot *Bot) {
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal(err)
	}

	bot = &Bot{dg}
	return
}

func (bot *Bot) Run() {
	bot.dg.AddHandler(bot.handleMessage)
	err := bot.dg.Open()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected")
}

func (bot *Bot) Close() {
	bot.dg.Close()
}

func (bot *Bot) handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Println(m.Author.Username, m.Content)
}
