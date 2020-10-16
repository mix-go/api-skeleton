package globals

import (
    "github.com/mix-go/console"
    "github.com/mix-go/logrus"
)

func Logger() *logrus.Logger {
    return console.Get("logger").(*logrus.Logger)
}
