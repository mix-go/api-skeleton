package commands

import (
    "fmt"
    "github.com/mix-go/gin"
    "github.com/mix-go/mix-api-skeleton/globals"
    "github.com/mix-go/mix-api-skeleton/routes/api"
)

type APICommand struct {
}

func (t *APICommand) Main() {
    // route
    router := gin.New(api.RouteDefinitions...)
    // logger
    router.Use(gin.LogrusWithFormatter(globals.Logger(), func(params gin.LogFormatterParams) string {
        return fmt.Sprintf("%s|%s|%d|%s",
            params.Method,
            params.Path,
            params.StatusCode,
            params.ClientIP,
        )
    }))
    // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
    if err := router.Run(); err != nil {
        panic(err)
    }
}
