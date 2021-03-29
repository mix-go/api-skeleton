package di

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/mix-go/xcli"
	"github.com/mix-go/xdi"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func init() {
	obj := xdi.Object{
		Name: "logrus",
		New: func() (i interface{}, e error) {
			logger := logrus.New()

			var fileRotate *rotatelogs.RotateLogs
			if err := xdi.Populate("file-rotatelogs", &fileRotate); err != nil {
				return nil, err
			}
			writer := io.MultiWriter(os.Stdout, fileRotate)
			logger.SetOutput(writer)
			if xcli.App().Debug {
				logger.SetLevel(logrus.DebugLevel)
			}

			return logger, nil
		},
		Singleton: true,
	}
	if err := xdi.Provide(&obj); err != nil {
		panic(err)
	}
}
