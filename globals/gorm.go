package globals

import (
	"github.com/mix-go/xdi"
	"gorm.io/gorm"
)

func Gorm() (db *gorm.DB) {
	if err := xdi.Populate("gorm", &db); err != nil {
		panic(err)
	}
	return
}
