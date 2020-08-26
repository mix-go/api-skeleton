package commands

import (
    "github.com/mix-go/mix-api-skeleton/commands"
    "github.com/mix-go/bean"
    "github.com/mix-go/console"
)

func init() {
    Commands = append(Commands,
        console.CommandDefinition{
            Name:  "hello",
            Usage: "\tEcho demo",
            Options: []console.OptionDefinition{
                {
                    Names: []string{"n", "name"},
                    Usage: "Your name",
                },
                {
                    Names: []string{"say"},
                    Usage: "\tSay ...",
                },
            },
            Reflect: bean.NewReflect(commands.HelloCommand{}),
        },
    )
}
