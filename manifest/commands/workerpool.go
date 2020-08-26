package commands

import (
    "github.com/mix-go/mix-api-skeleton/commands"
    "github.com/mix-go/bean"
    "github.com/mix-go/console"
)

func init() {
    Commands = append(Commands,
        console.CommandDefinition{
            Name:    "wp",
            Usage:   "\tWorker pool demo",
            Reflect: bean.NewReflect(commands.WorkerPoolCommand{}),
        },
        console.CommandDefinition{
            Name:  "wpd",
            Usage: "\tWorker pool daemon demo",
            Options: []console.OptionDefinition{
                {
                    Names: []string{"d", "daemon"},
                    Usage: "Run in the background",
                },
            },
            Reflect: bean.NewReflect(commands.WorkerPoolDaemonCommand{}),
        },
    )
}
