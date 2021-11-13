package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"
)

// chan读写特性中，当这个chan中数据为空时，receiver 接收数据的时候就会阻塞等待，直到 chan 被关闭或者有新的数据到来。

func doCleanup(closed chan struct{}) {
    time.Sleep((time.Minute))
    close(closed)
}

func main() {
    // closing，代表程序退出，但是清理工作还没做。
    var closing = make(chan struct{})
    // closed，代表清理工作已经做完。
    var closed = make(chan struct{})
    go func() {
        // 模拟业务处理
        for {
            select {
            case <-closing:
                return
            default:
                // ....... 业务计算
                time.Sleep(100 * time.Millisecond)
            }
        }
    }()
    // 处理CTRL+C等中断信号
    termChan := make(chan os.Signal)
    signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
    <-termChan
    close(closing)
    // 执行退出之前的清理动作
    go doCleanup(closed)
    select {
    case <-closed:
    case <-time.After(time.Second):
        fmt.Println("清理超时，不等了")
    }
    fmt.Println("优雅退出")
}
