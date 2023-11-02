package commands

import (
	"github.com/bwmarrin/discordgo"
)

type CommandHandler func(s *discordgo.Session, i *discordgo.InteractionCreate)

type BotCommands struct {
	Command *discordgo.ApplicationCommand
	Handler CommandHandler
}

var Commands = map[string]*BotCommands{
	"ask":     Ask,
	"grammar": Grammar,
}
