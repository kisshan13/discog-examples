package simple

import (
	discog "github.com/kisshan13/discog"
)

func GetGroup() *discog.CommandGroup {
	group := discog.NewCommandGroup("Simple", "Public commands for the Bot", "1076465065093513248")

	group.AddCommand(GetHelloCommand())

	return group
}
