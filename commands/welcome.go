package commands

import (
    "fmt"
    "github.com/mix-go/console"
    "github.com/mix-go/console/flag"
    "runtime"
    "strings"
)

const logo = `             ___         
 ______ ___  _ /__ ___ _____ ______ 
  / __ *__ \/ /\ \/ /__  __ */  __ \
 / / / / / / / /\ \/ _  /_/ // /_/ /
/_/ /_/ /_/_/ /_/\_\  \__, / \____/ 
                     /____/
`

func welcome() {
    fmt.Println(strings.Replace(logo, "*", "`", -1))
    fmt.Println("")
    fmt.Println(fmt.Sprintf("Server      Name:     %s", "mix-api"))
    fmt.Println(fmt.Sprintf("System      Name:     %s", runtime.GOOS))
    fmt.Println(fmt.Sprintf("Go          Version:  %s", runtime.Version()))
    fmt.Println(fmt.Sprintf("Framework   Version:  %s", console.App().AppVersion))
    fmt.Println(fmt.Sprintf("Listen      Addr:     %s", flag.Match("a", "addr").String(":8080")))
}
