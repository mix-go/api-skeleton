package commands

import (
    "fmt"
    "github.com/mix-go/console/flag"
)

type HelloCommand struct {
}

func (t *HelloCommand) Main() {
    name := flag.Match("n", "name").String("Xiao Ming")
    say := flag.Match("say").String("Hello, World!")
    fmt.Printf("%s: %s\n", name, say)
}
