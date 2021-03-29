package configor

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/mix-go/api-skeleton/globals"
	"github.com/mix-go/xcli/argv"
)

func init()  {
	// Conf support YAML, JSON, TOML, Shell Environment
	if err := configor.Load(&globals.Config, fmt.Sprintf("%s/../conf/config.yml", argv.Program().Dir)); err != nil {
		panic(err)
	}
}
