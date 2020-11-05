package commands

import (
    "context"
    "fmt"
    gin2 "github.com/gin-gonic/gin"
    "github.com/mix-go/console/flag"
    "github.com/mix-go/dotenv"
    "github.com/mix-go/gin"
    "github.com/mix-go/api-skeleton/globals"
    "github.com/mix-go/api-skeleton/routes/api"
    "net/http"
    "os"
    "os/signal"
    "strings"
    "syscall"
    "time"
)

const Addr = ":8080"

type APICommand struct {
}

func (t *APICommand) Main() {
    logger := globals.Logger()

    gin.SetMode(dotenv.Getenv("GIN_MODE").String(gin.ReleaseMode))

    // server
    router := gin.New(api.RouteDefinitionCallbacks...)
    srv := &http.Server{
        Addr:    flag.Match("a", "addr").String(Addr),
        Handler: router,
    }

    // signal
    ch := make(chan os.Signal)
    signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-ch
        logger.Info("Server shutdown")
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()
        if err := srv.Shutdown(ctx); err != nil {
            globals.Logger().Errorf("Server shutdown error: %s", err)
        }
    }()

    // error handle
    router.Use(gin2.Recovery())

    // logger
    router.Use(gin.LoggerWithFormatter(logger, func(params gin.LogFormatterParams) string {
        return fmt.Sprintf("%s|%s|%d|%s",
            params.Method,
            params.Path,
            params.StatusCode,
            params.ClientIP,
        )
    }))

    // run
    welcome()
    logger.Info("Server start")
    if err := srv.ListenAndServe(); err != nil && !strings.Contains(err.Error(), "http: Server closed") {
        panic(err)
    }
}
