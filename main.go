package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Set up run time flags
	botToken := flag.String("token", "", "Discord bot token.")
	logDebug := flag.Bool("debug", false, "Set logging level to debug.")

	flag.Parse()

	// Setup debug logging
	if *logDebug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// Check for required token
	if *botToken == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	session, err := discordgo.New(fmt.Sprintf("Bot %s", *botToken))
	if err != nil {
		log.Fatal(errors.WithMessage(err, "Error starting and connecting discord bot session"))
		os.Exit(1)
	}

	defer session.Close()

	router := dgc.Create(&dgc.Router{
		Prefixes: []string{"!"},
	})

	// Register a simple ping command
	router.RegisterCmd(&dgc.Command{
		Name:        "VibeCheck",
		Description: "Checks the vibe of a users.'",
		Usage:       "VibeCheck",
		Example:     "VibeCheck",
		IgnoreCase:  true,
		Handler: func(ctx *dgc.Ctx) {
			fmt.Printf("~~~~%+v\n", ctx.Event)
			fmt.Printf("~~~~%+v\n", ctx.Arguments)
			fmt.Printf("~~~~%+v\n", ctx.CustomObjects)
			fmt.Printf("~~~~%+v\n", ctx.Router)
			fmt.Printf("~~~~%+v\n", ctx.Command)
			ctx.RespondText("Pong!")
		},
	})

	// Initialize the router
	router.Initialize(session)

	err = session.Open()
	if err != nil {
		log.Fatal(errors.WithMessage(err, "Error opening Discord session"))
		os.Exit(1)
	}

	fmt.Println("Vibe Check is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
