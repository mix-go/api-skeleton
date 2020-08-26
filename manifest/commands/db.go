package commands

import (
    "github.com/mix-go/mix-api-skeleton/commands"
    "github.com/mix-go/bean"
    "github.com/mix-go/console"
)

func init() {
    Commands = append(Commands,
        console.CommandDefinition{
            Name:  "db",
            Usage: "\tDatabase query demo",
            Reflect: bean.NewReflect(commands.DBCommand{}),
        },
    )
}
