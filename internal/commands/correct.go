package commands

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cyneptic/cynbot/internal/services"
)

var Correct = &BotCommands{
	Command: &discordgo.ApplicationCommand{
		Name:        "correct",
		Description: "Check if something is grammatically correct and receive suggestions",
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

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "loading...",
			},
		})

		var service services.AskService
		service = services.NewAskService(fmt.Sprintf("is \"%s\" grammatically correct? what other ways can I write this", sentence))
		res, err := service.Process()
		if err != nil {
			log.Fatalf("error from service processor: %s", err.Error())
		}

		for err != nil && res != "" {
			time.Sleep(100 * time.Millisecond)
		}

		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &res,
		})
	},
}
