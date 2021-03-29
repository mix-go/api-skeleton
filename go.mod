module github.com/mix-go/api-skeleton

go 1.13

replace (
	github.com/mix-go/dotenv => ../mix/src/dotenv
	github.com/mix-go/xcli => ../mix/src/xcli
	github.com/mix-go/xdi => ../mix/src/xdi
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/go-redis/redis/v8 v8.8.0
	github.com/go-session/redis v3.0.1+incompatible
	github.com/go-session/session v3.1.2+incompatible
	github.com/golang/protobuf v1.5.1 // indirect
	github.com/jinzhu/configor v1.2.0
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.4 // indirect
	github.com/mix-go/dotenv v1.0.22
	github.com/mix-go/xcli v1.1.1
	github.com/mix-go/xdi v1.1.1
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/ugorji/go v1.2.4 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sys v0.0.0-20210326220804-49726bf1d181 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.6
)
