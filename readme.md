# Discog Example

Discog is a [Go](https://golang.org/) package and a discord bot library which uses [Discordgo](https://github.com/bwmarrin/discordgo) to interact with Discord API.

Discog provides a boilerplate less way to build discord bot. It has following features :

- Command manager
- Interaction Manager
- Handlers

**For help with this package or general Go discussion, please join the [Discog Discord Server](https://discord.gg/golang) chat server.**

## Getting Started

### Installing

Discog uses **Discordgo** so it's important to upto date with **Discordgo** version for **Discog** which is `v0.28.1`

`go get` _will always pull the latest tagged release from the master branch._

```sh
go get github.com/kisshan13/discog
```

### Basic Structure

### A typical top-level directory layout

    .
    ├── command                 # Commands Folder
	│   ├── simple          	# Command Info
	│   │	├── commands.go		# Slash command info 
	│   │	├── group.go 		# Command group
	│   └── manager.go          # Command manager
    ├── go.mod     
    ├── go.sum                  
    └── main.go                 # Entry point

`main.go`

```go
package main

import (
	"log"

	commands "discog-examples/command"

	"github.com/bwmarrin/discordgo"
	discog "github.com/kisshan13/discog"
)

func main() {

	manager := commands.GetManager()
	bot, err := discog.NewBot("<bot-token>", manager)

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

```