package main

import (
    "context"
    "fmt"
    "time"
)

// Go语言context标准库的Context类型提供了一个Done()方法，该方法返回一个类型为<-chan struct{}的channel。
// 每次context收到取消事件后这个channel都会接收到一个struct{}类型的值。
// 所以在Go语言里监听取消事件就是等待接收<-ctx.Done()。

func main() {
    ch := make(chan struct{})
    ctx, cancel := context.WithCancel(context.Background())

    go func(ctx context.Context) {
        for {
            select {
            case <-ctx.Done():
                ch <- struct{}{}
                return
            default:
                fmt.Println("Hello...")
            }

            time.Sleep(500 * time.Millisecond)
        }
    }(ctx)

    go func() {
        time.Sleep(3 * time.Second)
        cancel()
    }()

    <-ch
    fmt.Println("结束")
}
