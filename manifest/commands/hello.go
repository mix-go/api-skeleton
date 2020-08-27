package commands

import (
    "github.com/mix-go/console"
    "github.com/mix-go/mix-api-skeleton/commands"
)

func init() {
    Commands = append(Commands,
        console.CommandDefinition{
            Name:  "api",
            Usage: "\tStart the api server",
            Options: []console.OptionDefinition{
                {
                    Names: []string{"h", "host"},
                    Usage: "\tListen to the specified host",
                },
                {
                    Names: []string{"p", "port"},
                    Usage: "\tListen to the specified port",
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
