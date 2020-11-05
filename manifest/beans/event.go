package beans

import (
    "github.com/mix-go/bean"
    "github.com/mix-go/event"
    "github.com/mix-go/api-skeleton/listeners"
)

func Event() {
    Beans = append(Beans,
        bean.BeanDefinition{
            Name:            "eventDispatcher",
            Reflect:         bean.NewReflect(event.NewDispatcher),
            Scope:           bean.SINGLETON,
            ConstructorArgs: bean.ConstructorArgs{&listeners.CommandListener{}},
        },
    )
}
