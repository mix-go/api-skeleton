package main

import (
    "github.com/mix-go/mix-api-skeleton/globals"
    "github.com/mix-go/mix-api-skeleton/manifest"
    "fmt"
    "github.com/jinzhu/configor"
    console2 "github.com/mix-go/console"
    "github.com/mix-go/console/argv"
    "github.com/mix-go/dotenv"
)

func init() {
    // Env
    if err := dotenv.Load(fmt.Sprintf("%s/../.env", argv.Program().Dir)); err != nil {
        panic(err)
    }
    // Conf support YAML, JSON, TOML, Shell Environment
    if err := configor.Load(&globals.Config, fmt.Sprintf("%s/../conf/config.json", argv.Program().Dir)); err != nil {
        panic(err)
    }
    // Manifest
    manifest.Init()
}

func main() {
    // App
    console2.NewApplication(manifest.ApplicationDefinition, "eventDispatcher", "error").Run()
}
