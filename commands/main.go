package commands

import (
	"github.com/mix-go/xcli"
)

var Commands = []*xcli.Command{
	{
		Name:  "api",
		Short: "\tStart the api server",
		Options: []*xcli.Option{
			{
				Names: []string{"a", "addr"},
				Short: "\tListen to the specified address",
			},
			{
				Names: []string{"d", "daemon"},
				Short: "\tRun in the background",
			},
		},
		RunI: &APICommand{},
	},
}
