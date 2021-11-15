package main

import (
    "fmt"
    "sync"
)

// 结果： &{{0 0} 张三 0}
type user struct {
    lock sync.Mutex
    name string
    age  int
}

func main() {

    u := new(user) // 默认给u分配到内存全部为0

    u.lock.Lock() // 可以直接使用，因为lock为0,是开锁状态
    u.name = "张三"
    u.lock.Unlock()

    fmt.Println(u)
}
