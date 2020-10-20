package main

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	PREFIX         = "self/"
	EMOJI_THUMBSUP = "\xf0\x9f\x91\x8d"
	EMOJI_NO       = "\xf0\x9f\x9a\xab"
)

type Bot struct {
	dg *discordgo.Session
}

func newBot(token string) (bot *Bot) {
	dg, err := discordgo.New("Bot " + token)
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

	parts := strings.Split(m.Content, " ")
	if len(parts) < 2 {
		return
	}

	if parts[0] == "!role" {
		switch parts[1] {
		case "get":
			if len(parts) < 3 {
				return
			}
			bot.getRole(s, m, parts[2])
		case "remove":
			if len(parts) < 3 {
				return
			}
			bot.removeRole(s, m, parts[2])
		case "list":
			bot.listRoles(s, m)
		case "source":
			s.ChannelMessageSend(m.ChannelID, "https://github.com/acmumn/rolebot")
		}
	}
}

func (bot *Bot) getRole(s *discordgo.Session, m *discordgo.MessageCreate, rolename string) (err error) {
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		return err
	}
	roles, err := s.GuildRoles(channel.GuildID)
	if err != nil {
		return err
	}
	found := false
	for _, role := range roles {
		if !strings.HasPrefix(role.Name, PREFIX) {
			continue
		}
		if (PREFIX+rolename) == role.Name || rolename == role.Name {
			err = s.GuildMemberRoleAdd(channel.GuildID, m.Author.ID, role.ID)
			log.Println(err)
			if err != nil {
				break
			}
			found = true
		}
	}
	var emoji string
	if found {
		emoji = EMOJI_THUMBSUP
	} else {
		emoji = EMOJI_NO
	}
	err = s.MessageReactionAdd(m.ChannelID, m.ID, emoji)
	if err != nil {
		return err
	}
	return nil
}

func (bot *Bot) removeRole(s *discordgo.Session, m *discordgo.MessageCreate, rolename string) (err error) {
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		return err
	}
	roles, err := s.GuildRoles(channel.GuildID)
	if err != nil {
		return err
	}
	found := false
	for _, role := range roles {
		if !strings.HasPrefix(role.Name, PREFIX) {
			continue
		}
		if (PREFIX+rolename) == role.Name || rolename == role.Name {
			err = s.GuildMemberRoleRemove(channel.GuildID, m.Author.ID, role.ID)
			if err != nil {
				break
			}
			found = true
		}
	}
	var emoji string
	if found {
		emoji = EMOJI_THUMBSUP
	} else {
		emoji = EMOJI_NO
	}
	err = s.MessageReactionAdd(m.ChannelID, m.ID, emoji)
	if err != nil {
		return err
	}
	return nil
}

func (bot *Bot) listRoles(s *discordgo.Session, m *discordgo.MessageCreate) (err error) {
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		return err
	}
	roles, err := s.GuildRoles(channel.GuildID)
	if err != nil {
		return err
	}
	list := make([]string, 0)
	for _, role := range roles {
		if !strings.HasPrefix(role.Name, PREFIX) {
			continue
		}
		list = append(list, "`"+role.Name+"`")
	}
	s.ChannelMessageSend(m.ChannelID, strings.Join(list, ", "))
	return nil
}
