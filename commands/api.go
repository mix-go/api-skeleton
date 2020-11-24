package commands

import (
    "context"
    "fmt"
    "github.com/mix-go/api-skeleton/globals"
    "github.com/mix-go/api-skeleton/routes"
    "github.com/mix-go/console/flag"
    "github.com/mix-go/dotenv"
    "github.com/mix-go/gin"
    "net/http"
    "os"
    "os/signal"
    "strings"
    "syscall"
    "time"
)

const Addr = ":8080"

var Server *http.Server

type APICommand struct {
}

func (t *APICommand) Main() {
    logger := globals.Logger()
    mode := dotenv.Getenv("GIN_MODE").String(gin.ReleaseMode)

    // server
    gin.SetMode(mode)
    router := gin.New(routes.ApiDefinition())
    srv := &http.Server{
        Addr:    flag.Match("a", "addr").String(Addr),
        Handler: router,
    }
    Server = srv

    // signal
    ch := make(chan os.Signal)
    signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-ch
        logger.Info("Server shutdown")
        ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
        if err := srv.Shutdown(ctx); err != nil {
            logger.Errorf("Server shutdown error: %s", err)
        }
    }()

    // logger
    if mode != gin.ReleaseMode {
        router.Use(gin.LoggerWithFormatter(logger, func(params gin.LogFormatterParams) string {
            return fmt.Sprintf("%s|%s|%d|%s",
                params.Method,
                params.Path,
                params.StatusCode,
                params.ClientIP,
            )
        }))
    }

    // run
    welcome()
    logger.Info("Server start")
    if err := srv.ListenAndServe(); err != nil && !strings.Contains(err.Error(), "http: Server closed") {
        panic(err)
    }
}
