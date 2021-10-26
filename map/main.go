package main

import (
    "fmt"
    "sync"
)

// 解析：不能编译，sync.Map 没有 Len 方法
func main() {
    var m sync.Map
    m.LoadOrStore("a", 1)
    m.Delete("a")
    fmt.Println(m.Len())
}
