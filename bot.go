package main

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
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
	if strings.HasPrefix(m.Content, "!role") {
		parts := strings.Split(m.Content, " ")
		log.Println(parts)

		if len(parts) < 1 {
			return
		}
		switch parts[1] {
		case "get":
			if len(parts) < 2 {
				return
			}
			bot.getRole(s, m, parts[2])
		case "remove":
			if len(parts) < 2 {
				return
			}
			bot.removeRole(s, m, parts[2])
		case "list":
			bot.listRoles(s, m)
		}
	} else if strings.HasPrefix(m.Content, "!source") {
		s.ChannelMessageSend(m.ChannelID, "https://github.com/acmumn/rolebot")
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
		if !strings.HasPrefix(role.Name, "self/") {
			continue
		}
		if ("self/" + rolename) == role.Name {
			err = s.GuildMemberRoleAdd(channel.GuildID, m.Author.ID, role.ID)
			if err != nil {
				break
			}
			found = true
		}
	}
	var emoji string
	if found {
		emoji = "\xf0\x9f\x91\x8d"
	} else {
		emoji = "\xf0\x9f\x9a\xab"
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
		if !strings.HasPrefix(role.Name, "self/") {
			continue
		}
		if ("self/" + rolename) == role.Name {
			err = s.GuildMemberRoleRemove(channel.GuildID, m.Author.ID, role.ID)
			if err != nil {
				break
			}
			found = true
		}
	}
	var emoji string
	if found {
		emoji = "\xf0\x9f\x91\x8d"
	} else {
		emoji = "\xf0\x9f\x9a\xab"
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
		if !strings.HasPrefix(role.Name, "self/") {
			continue
		}
		list = append(list, "`"+role.Name+"`")
	}
	s.ChannelMessageSend(m.ChannelID, strings.Join(list, ", "))
	return nil
}
