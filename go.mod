module github.com/mix-go/mix-api-skeleton

go 1.13

replace (
	github.com/mix-go/bean => ../mix/src/bean
	github.com/mix-go/console => ../mix/src/console
	github.com/mix-go/dotenv => ../mix/src/dotenv
	github.com/mix-go/event => ../mix/src/event
	github.com/mix-go/logrus => ../mix/src/logrus
	github.com/mix-go/workerpool => ../mix/src/workerpool
)

require (
	github.com/go-redis/redis/v8 v8.0.0-beta.7
	github.com/jinzhu/configor v1.2.0 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/mix-go/bean v1.0.1
	github.com/mix-go/console v1.0.1
	github.com/mix-go/dotenv v1.0.1
	github.com/mix-go/event v1.0.1
	github.com/mix-go/logrus v1.0.1
	github.com/mix-go/workerpool v1.0.1
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gorm.io/gorm v0.2.34 // indirect
)
