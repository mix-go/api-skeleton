package manifest

import (
    "github.com/mix-go/mix-api-skeleton/manifest/beans"
    "github.com/mix-go/mix-api-skeleton/manifest/commands"
    "github.com/mix-go/console"
    "github.com/mix-go/dotenv"
)

var (
    ApplicationDefinition console.ApplicationDefinition
)

func Init() {
    beans.Init()

    ApplicationDefinition = console.ApplicationDefinition{
        AppName:    "app",
        AppVersion: "1.0.0-alpha",
        AppDebug:   dotenv.Getenv("APP_DEBUG").Bool(),
        Beans:      beans.Beans,
        Commands:   commands.Commands,
    }
}
