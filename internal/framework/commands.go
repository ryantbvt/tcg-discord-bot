package framework

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type CommandHandler interface {
	// Definition returns ApplicationCommand sent to Discord
	Definition() *discordgo.ApplicationCommand
	// Handle is called when the command is invoked
	Handle(s *discordgo.Session, i *discordgo.InteractionCreate)
	// serverless/lambda support
	HandleSync(i *discordgo.Interaction) *discordgo.InteractionResponse
}

type CommandRegistry struct {
	handlers map[string]CommandHandler
	order    []string
}

func newCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		handlers: make(map[string]CommandHandler),
	}
}

// register a CommandHandler
func (r *CommandRegistry) Add(h CommandHandler) {
	name := h.Definition().Name

	if _, exists := r.handlers[name]; exists {
		log.Panicf("discord: duplicate command registered: %q", name)
	}

	r.handlers[name] = h
	r.order = append(r.order, name)
}

// checks for commands w/ no handler
func (r *CommandRegistry) dispatch(s *discordgo.Session, i *discordgo.InteractionCreate) {
	name := i.ApplicationCommandData().Name
	h, ok := r.handlers[name]
	if !ok {
		log.Printf("discord: no handler for command: %q", name)
		return
	}
	h.Handle(s, i)
}

// syncs commands to Discord
func (r *CommandRegistry) sync(s *discordgo.Session) error {
	appID := s.State.User.ID
	guildID := "" // replace guild ID to test command instantly in a single server

	defs := make([]*discordgo.ApplicationCommand, 0, len(r.order))
	for _, name := range r.order {
		defs = append(defs, r.handlers[name].Definition())
	}

	for _, def := range defs {
		if _, err := s.ApplicationCommandCreate(appID, guildID, def); err != nil {
			return fmt.Errorf("discord: failed to register command %q: %w", def.Name, err)
		}
		log.Printf("discord: registered command: %q", def.Name)
	}

	return nil
}

func (r *CommandRegistry) dispatchSync(i *discordgo.Interaction) *discordgo.InteractionResponse {
	name := i.ApplicationCommandData().Name
	h, ok := r.handlers[name]

	if !ok {
		log.Printf("discord: no handler for command %q", name)

		return &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Error: Command not found",
			},
		}
	}
	return h.HandleSync(i)
}
