package main

import "sync"

// 知识点：go 内建的 map 对象不是线程(goroutine)安全的，并发读写的时候运行时会有检查， 遇到并发问题就会导致 panic。
// 结果：fatal error: concurrent map read and map write

func main() {
    var mu sync.Mutex
    var m = make(map[int]int, 10) // 初始化一个map
    go func() {
        for {
            mu.Lock()
            m[1] = 1 // 设置key
            mu.Unlock()
        }
    }()

    go func() {
        for {
            mu.Lock()
            _ = m[2] // 访问这个map
            mu.Unlock()
        }
    }()
    select {}
}
