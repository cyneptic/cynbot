package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cyneptic/cynbot/grammar"
)

var Ask = &BotCommands{
	Command: &discordgo.ApplicationCommand{
		Name:        "bitch",
		Description: "Ask a query from bitch",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "query",
				Description: "query that has been asked of bitch",
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
