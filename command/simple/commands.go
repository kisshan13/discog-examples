package simple

import (
	"github.com/bwmarrin/discordgo"
	discog "github.com/kisshan13/discog"
)

func GetHelloCommand() *discog.SlashCommand {
	command := discog.NewSlashCommand(discordgo.ApplicationCommand{
		Name:        "hello",
		Description: "Replies with a Hii message",
	})

	command.SetHandler(func(ctx *discog.Ctx) {

		ctx.SetContent("Hii")
		err := ctx.Send()
		if err != nil {
			println(err.Error())
		}
	})

	return command
}
