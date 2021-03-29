package globals

import (
	"github.com/mix-go/xdi"
	"github.com/sirupsen/logrus"
)

func Logrus() (logger *logrus.Logger) {
	if err := xdi.Populate("logrus", &logger); err != nil {
		panic(err)
	}
	return
}
