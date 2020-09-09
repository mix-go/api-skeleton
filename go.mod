module github.com/mix-go/mix-api-skeleton

go 1.13

replace (
	github.com/mix-go/bean => ../mix/src/bean
	github.com/mix-go/console => ../mix/src/console
	github.com/mix-go/dotenv => ../mix/src/dotenv
	github.com/mix-go/event => ../mix/src/event
	github.com/mix-go/gin => ../mix/src/gin
	github.com/mix-go/logrus => ../mix/src/logrus
	github.com/mix-go/workerpool => ../mix/src/workerpool
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v8 v8.0.0-beta.7
	github.com/jinzhu/configor v1.2.0
	github.com/jinzhu/gorm v1.9.16
	github.com/mix-go/bean v1.0.8
	github.com/mix-go/console v1.0.9
	github.com/mix-go/dotenv v1.0.1
	github.com/mix-go/event v1.0.1
	github.com/mix-go/gin v1.0.9
	github.com/mix-go/logrus v1.0.8
	github.com/mix-go/workerpool v1.0.8
	golang.org/x/sys v0.0.0-20200826173525-f9321e4c35a6 // indirect
	gorm.io/gorm v0.2.34 // indirect
)
