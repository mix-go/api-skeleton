package globals

import (
    "fmt"
    "github.com/mix-go/console"
    "github.com/mix-go/logrus"
    "io"
    "os"
)

func Init() {
    // 日志配置
    logger := Logger()
    file := logrus.NewFileWriter(fmt.Sprintf("%s/../runtime/logs/test.log", console.App().BasePath), 7)
    writer := io.MultiWriter(os.Stdout, file)
    logger.SetOutput(writer)
}
