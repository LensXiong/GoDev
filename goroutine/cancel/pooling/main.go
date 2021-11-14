package main

import (
    "fmt"
    "time"
)
// 解析：
// 结果：
// 接收到的值:  Hello
// 接收到的值:  Hello
// 接收到的值:  Hello
// 结束
func main() {
    ch := make(chan string, 6)
    done := make(chan struct{})
    go func() {
        for {
            select {
            case ch <- "Hello":
            case <-done:
                close(ch)
                return
            }
            time.Sleep(1 * time.Second)
        }
    }()

    go func() {
        time.Sleep(3 * time.Second)
        done <- struct{}{}
    }()

    for i := range ch {
        fmt.Println("接收到的值: ", i)
    }

    fmt.Println("结束")
}
