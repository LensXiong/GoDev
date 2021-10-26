package main

import "fmt"

// 结果：① panic: close of nil channel ② panic: send on closed channel
// 解析：两个问题：① ch 未有被初始化；② 关闭时会报错。
// channel 是引用类型。
// channel 必须初始化才能写入数据, 即 make 后才能使用，否则关闭时会报错。intChan = make(chan int, 3)
// 使用内置函数 close 可以关闭 channel, 当 channel 关闭后，就不能再向 channel 写数据了，但是仍然可以从该 channel 读取数据。
func main() {
    var ch chan int
    // ① 需要初始化 channel
    // ch = make(chan int, 1)
    var count int
    go func() {
        ch <- 1
    }()
    go func() {
        count++
        close(ch)
    }()
    <-ch
    // ② 应在此处关闭 channel
    // close(ch)
    fmt.Println(count)
}
