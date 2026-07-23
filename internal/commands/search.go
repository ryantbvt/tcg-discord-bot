package commands

import (
	"errors"
	"log"

	"github.com/bwmarrin/discordgo"

	"github.com/ryantbvt/tcg-discord-bot/internal/tcgapi"
)

type SingleSearchCommand struct {
	client *tcgapi.Client
}

func NewSingleSearchCommand(client *tcgapi.Client) *SingleSearchCommand {
	return &SingleSearchCommand{client: client}
}

func (c *SingleSearchCommand) Definition() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "searchbyid",
		Description: "Returns card information, searched by id",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "Card ID, e.g. swsh3-136",
				Required:    true,
			},
		},
	}
}

func (c *SingleSearchCommand) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	id := i.ApplicationCommandData().Options[0].StringValue()

	card, err := c.client.GetCard(id)
	if err != nil {
		msg := "Something went wrong look that up."
		if errors.Is(err, tcgapi.ErrNotFound) {
			msg = "Couldn't find a card with that ID."
		} else {
			log.Printf("SingleSearchCommand: GetCard(%q) failed: %v", id, err)
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: msg,
			},
		})

		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: card.Name + " - " + card.Rarity,
		},
	})
}

func (c *SingleSearchCommand) HandleSync(i *discordgo.Interaction) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	}
}
