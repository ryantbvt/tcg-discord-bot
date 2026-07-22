package framework

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

// Router holds all registered handlers and dispatches incoming interactions
type Router struct {
	commands *CommandRegistry
	// Add buttons/modals if needed
}

func NewRouter() *Router {
	return &Router{
		commands: newCommandRegistry(),
		// Add buttons/modals if needed
	}
}

func (r *Router) Commands() *CommandRegistry {
	return r.commands
}

// can create similar functions as above for buttons/modals

// used for normal bot / long living sessions
func (r *Router) Handler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			r.commands.dispatch(s, i)

		default:
			log.Printf("unhandled interaction type: %v", i.Type)

		}
	}
}

func (r *Router) Sync(s *discordgo.Session) error {
	return r.commands.sync(s)
}

// lambda specific handler - serverless
func (r *Router) HandleLambda(body string) (string, error) {
	var i discordgo.Interaction
	if err := json.Unmarshal([]byte(body), &i); err != nil {
		return "", err
	}

	if i.Type == discordgo.InteractionPing {
		return `{"type": 1}`, nil
	}

	// Handle command
	if i.Type == discordgo.InteractionApplicationCommand {
		resp := r.commands.dispatchSync(&i)
		respJson, _ := json.Marshal(resp)

		return string(respJson), nil
	}

	return "", fmt.Errorf("unhandled interaction type")
}
