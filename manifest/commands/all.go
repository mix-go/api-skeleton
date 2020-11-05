package commands

import (
    "github.com/mix-go/api-skeleton/commands"
    "github.com/mix-go/console"
)

var (
    Commands []console.CommandDefinition
)

func init() {
    Commands = append(Commands,
        console.CommandDefinition{
            Name:  "api",
            Usage: "\tStart the api server",
            Options: []console.OptionDefinition{
                {
                    Names: []string{"a", "addr"},
                    Usage: "\tListen to the specified address",
                },
                {
                    Names: []string{"d", "daemon"},
                    Usage: "\tRun in the background",
                },
            },
            Command: &commands.APICommand{},
        },
    )
}
