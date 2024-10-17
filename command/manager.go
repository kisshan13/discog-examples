package commands

import (
	simple "discog-examples/command/simple"

	discog "github.com/kisshan13/discog"
)

func GetManager() *discog.CommandManager {
	manager := discog.NewCommandManager()

	manager.AddGroup(simple.GetGroup())

	return manager
}
