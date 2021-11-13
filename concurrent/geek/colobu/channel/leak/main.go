package main

import (
    "fmt"
    "time"
)

// 知识点：关于 goroutine 泄漏，下面代码有什么问题？

// process 函数会启动一个 goroutine，去处理需要长时间处理的业务，处理完之后，会发送 true 到 chan 中，
// 目的是通知其它等待的 goroutine，可以继续处理了。主 goroutine 接收到任务处理完成的通知，或者超时后就返回了。这段代码有问题吗?
// 如果发生超时，process 函数就返回了，这就会导致 unbuffered 的 chan 从来就没有被读取。
// unbuffered chan 必须等 reader 和 writer 都准备好了才能交流，否则就会阻塞。
// 超时导致未读，结果就是子 goroutine 就阻塞在写永远结束不了，进而导致 goroutine 泄漏。
// 解决这个 Bug 的办法就是将 unbuffered chan 改成容量为 1 的 chan，这样写就不会被阻塞了。

func process(timeout time.Duration) bool {
    ch := make(chan bool, 1)
    go func() {
        // 模拟处理耗时的业务
        // time.Sleep((timeout + time.Second))
        ch <- true // block
        fmt.Println("exit goroutine")
    }()
    select {
    case result := <-ch:
        return result
    case <-time.After(timeout):
        return false
    }
}

func main() {
    res := process(1 * time.Second)
    fmt.Println(res)
}
