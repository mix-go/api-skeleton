package globals

import (
    "github.com/mix-go/console"
    "github.com/mix-go/event"
)

func EventDispatcher() *event.Dispatcher {
    return console.Get("eventDispatcher").(*event.Dispatcher)
}
