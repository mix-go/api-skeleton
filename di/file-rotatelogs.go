package di

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/mix-go/xcli"
	"github.com/mix-go/xdi"
)

func init() {
	obj := xdi.Object{
		Name: "file-rotatelogs",
		New: func() (i interface{}, e error) {
			filename := fmt.Sprintf("%s/../logs/cli.log", xcli.App().BasePath)
			options := []rotatelogs.Option{
				rotatelogs.WithLinkName(filename),
				rotatelogs.WithMaxAge(-1),
			}
			options = append(options, rotatelogs.WithRotationCount(7))
			return rotatelogs.New(filename+".%Y%m%d", options...)
		},
	}
	if err := xdi.Provide(&obj); err != nil {
		panic(err)
	}
}
