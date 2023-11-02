package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/cyneptic/cynbot/internal/services"
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
		sentence := i.ApplicationCommandData().Options[0].Value.(string)
		service := services.NewAskService(sentence)
		res, err := service.Process()
		if err != nil {
			log.Printf("error from service processor: %s", err.Error())
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: res,
			},
		})
	},
}
