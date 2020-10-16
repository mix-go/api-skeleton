package globals

import (
    "github.com/jinzhu/gorm"
    "github.com/mix-go/console"
)

func DB() *gorm.DB {
    return console.Get("db").(*gorm.DB)
}
