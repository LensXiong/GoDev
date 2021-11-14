package main

import (
    "fmt"
    "net/http"
    "os"
    "time"
)
// Go语言context标准库的Context类型提供了一个Done()方法，该方法返回一个类型为<-chan struct{}的channel。
// 每次context收到取消事件后这个channel都会接收到一个struct{}类型的值。
// 所以在Go语言里监听取消事件就是等待接收<-ctx.Done()。

func main() {
    // 创建一个监听8000端口的服务器
    http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()
        // 输出到STDOUT展示处理已经开始
        fmt.Fprint(os.Stdout, "processing request\n")
        // 通过select监听多个channel
        select {
        case <-time.After(2 * time.Second):
            // 如果两秒后接受到了一个消息后，意味请求已经处理完成
            // 我们写入"request processed"作为响应
            w.Write([]byte("request processed"))
        case <-ctx.Done():

            // 如果处理完成前取消了，在STDERR中记录请求被取消的消息
            fmt.Fprint(os.Stderr, "request cancelled\n")
        }
    }))
}
