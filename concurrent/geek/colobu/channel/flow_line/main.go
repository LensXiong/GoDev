package main

import (
    "fmt"
    "time"
)

// 使用 Channel 进行任务编排。
// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，
// 要求编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、......的顺序打印出来。

// 解析：注(i+1)%count 的循环。

// //首先把令牌交给第一个worker
type Token struct{}

func newWorker(id int, ch chan Token, nextCh chan Token) {
    for {
        token := <-ch
        fmt.Println(id + 1)
        time.Sleep(time.Second)
        nextCh <- token
    }
}

func main() {
    chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}
    count := 4
    // 创建n个worker
    for i := 0; i < count; i++ {
        go newWorker(i, chs[i], chs[(i+1)%count])
    }
    // 首先把令牌交给第一个worker
    chs[0] <- struct{}{}
    select {}
}
