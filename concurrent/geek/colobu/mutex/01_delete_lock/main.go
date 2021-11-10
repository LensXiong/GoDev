package main

import (
    "fmt"
    "sync"
)

// 使用 Mutex 常见的错误场景1： Lock/Unlock 不是成对出现。

// 缺少 Unlock 的场景：① 代码中有太多的 if-else 分支，可能在某个分支中漏写了 Unlock;
// ② 在重构的时候把 Unlock 给删除了。
// ③ Unlock 误写成了 Lock。

// 缺少 Lock 的场景：修改BUG或者重构时将 Lock 调用给删除了，或者注释掉了（切记）。

// 结果：fatal error: sync: unlock of unlocked mutex

// 解析：mu.Lock() 一行代码被删除了，直接 Unlock 一个未加锁的 Mutex 会 panic
func main() {
    var mu sync.Mutex
    defer mu.Unlock()
    fmt.Println("hello world!")
}
