package beans

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/mix-go/bean"
    "github.com/mix-go/dotenv"
)

func DB() {
    Beans = append(Beans,
        bean.BeanDefinition{
            Name:            "db",
            Reflect:         bean.NewReflect(gorm.Open),
            Scope:           bean.SINGLETON,
            ConstructorArgs: bean.ConstructorArgs{"mysql", dotenv.Getenv("DATABASE_DSN").String()},
        },
    )
}
