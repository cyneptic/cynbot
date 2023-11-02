package bot

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/cyneptic/cynbot/internal/commands"
)

type Bot struct {
	S                  *discordgo.Session
	Token              string
	GuildID            string
	RegisteredCommands []*discordgo.ApplicationCommand
}

func CreateNewBot(token, guildID string) *Bot {
	sesh, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Invalid bot parameter: %v", err)
	}

	sesh.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Connected as ", s.State.User.ID)
	})

	sesh.Open()

	b := &Bot{
		S:                  sesh,
		Token:              token,
		GuildID:            guildID,
		RegisteredCommands: make([]*discordgo.ApplicationCommand, len(commands.Commands)),
	}

	b.RegisterCommands()
	b.RegisterHandlers()

	return b
}

func (b *Bot) RegisterCommands() {
	for _, cmd := range commands.Commands {
		c, err := b.S.ApplicationCommandCreate(b.S.State.User.ID, b.GuildID, cmd.Command)
		if err != nil {
			log.Panicf("Adding command %v, got error %v", cmd.Command.Name, err)
		}
		b.RegisteredCommands = append(b.RegisteredCommands, c)
	}
}

func (b *Bot) RegisterHandlers() {
	b.S.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commands.Commands[i.ApplicationCommandData().Name]; ok {
			h.Handler(s, i)
		}
	})
}

func (b *Bot) RemoveCommands() {
	for _, v := range b.RegisteredCommands {
		err := b.S.ApplicationCommandDelete(b.S.State.User.ID, b.GuildID, v.ID)
		if err != nil {
			log.Panicf("cannot delete '%v' command: %v", v.Name, err)
		}
	}
}
