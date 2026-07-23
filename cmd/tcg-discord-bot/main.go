package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/ryantbvt/tcg-discord-bot/internal/commands"
	"github.com/ryantbvt/tcg-discord-bot/internal/framework"
	"github.com/ryantbvt/tcg-discord-bot/internal/tcgapi"
)

func main() {
	// load configs
	cfg := framework.LoadEnv()

	// TCGdex Client
	tcgClient := tcgapi.NewClient("en")

	// initialize discord bot
	svr, err := discordgo.New("Bot " + cfg.DiscToken)
	if err != nil {
		log.Fatal("Discord token not set")
	}

	// add router
	router := framework.NewRouter()

	// Slash commands
	router.Commands().Add(&commands.PingCommand{})
	router.Commands().Add(commands.NewSingleSearchCommand(tcgClient))

	// add handlers
	svr.AddHandler(router.Handler())
	svr.Identify.Intents = discordgo.IntentGuilds

	// open discord
	if err := svr.Open(); err != nil {
		log.Fatal("Error opening Discord connection", err)
	}

	defer svr.Close()

	// register slash commands with Discord
	if err := router.Sync(svr); err != nil {
		log.Fatalf("Error syncing commands: %v", err)
	}

	log.Println("Bot is now running")

	// wait for ctrl + C or sig
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	log.Println("Shutting down bot")
}
