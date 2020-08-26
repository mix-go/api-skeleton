package commands

import (
    "github.com/mix-go/mix-api-skeleton/globals"
    "context"
    "fmt"
    "github.com/mix-go/workerpool"
    "os"
    "os/signal"
    "strings"
    "syscall"
    "time"
)

type worker struct {
    workerpool.WorkerTrait
}

func (t *worker) Handle(data interface{}) {
    fmt.Println(data)
}

func NewWorker() workerpool.Worker {
    return &worker{}
}

type WorkerPoolCommand struct {
}

func (t *WorkerPoolCommand) Main() {
    jobQueue := make(chan interface{}, 50)
    d := workerpool.NewDispatcher(NewWorker, 15, jobQueue)

    go func() {
        for i := 0; i < 10000; i++ {
            jobQueue <- i
        }
        d.Stop()
    }()

    d.Run()
    d.Wait()
}

type WorkerPoolDaemonCommand struct {
}

func (t *WorkerPoolDaemonCommand) Main() {
    redis := globals.Redis()
    jobQueue := make(chan interface{}, 50)
    d := workerpool.NewDispatcher(NewWorker, 15, jobQueue)

    ch := make(chan os.Signal)
    signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-ch
        d.Stop()
    }()

    go func() {
        for {
            res, err := redis.BRPop(context.Background(), 3*time.Second, "test").Result()
            if err != nil {
                if strings.Contains(err.Error(), "redis: nil") {
                    continue
                }
                fmt.Println(fmt.Sprintf("Redis Error: %s", err))
                d.Stop();
                return
            }
            // brPop命令最后一个键才是值
            jobQueue <- res[1]
        }
    }()

    d.Run()
    d.Wait()
}
