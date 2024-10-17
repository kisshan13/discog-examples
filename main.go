package main

import (
	"log"

	commands "discog-examples/command"

	"github.com/bwmarrin/discordgo"
	discog "github.com/kisshan13/discog"
)

func main() {

	manager := commands.GetManager()
	bot, err := discog.NewBot("<token>", manager)

	if err != nil {
		panic(err)
	}

	bot.OnReady(func(r *discordgo.Ready) {
		log.Printf("Bot has connected to Discord as %s (%s)\n", r.User.Username, r.User.ID)
	})

	bot.Run(func(session *discordgo.Session, err error) {
		if err != nil {
			panic(err)
		}
		log.Println("Bot is running")
	})
}
