package globals

import (
    "fmt"
    "github.com/mix-go/console"
    "github.com/mix-go/logrus"
    logrus2 "github.com/sirupsen/logrus"
    "io"
    "os"
)

func Init() {
    // logger
    logger := Logger()
    file := logrus.NewFileWriter(fmt.Sprintf("%s/../runtime/logs/mix.log", console.App.BasePath), 7)
    writer := io.MultiWriter(os.Stdout, file)
    logger.SetOutput(writer)
    if console.App.AppDebug {
        logger.SetLevel(logrus2.DebugLevel)
    }

    // db
    db := DB()
    db.SetLogger(logger)
    db.LogMode(console.App.AppDebug)
}
