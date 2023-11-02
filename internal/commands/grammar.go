package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cyneptic/cynbot/grammar"
)

var Grammar = &BotCommands{
	Command: &discordgo.ApplicationCommand{
		Name:        "grammar",
		Description: "Grammar check a sentence.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "sentence",
				Description: "sentence to be checked",
				Required:    true,
			},
		},
	},
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// wrong, correct, err := grammar.Check(i.Message.Interaction)
		sentence := i.ApplicationCommandData().Options[0].Value.(string)
		var g grammar.GrammarInterface
		g = grammar.NewGrammarChecker(sentence)

		g.Check(sentence)

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Basic Hellow",
			},
		})
	},
}
