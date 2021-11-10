package main

import (
    "fmt"
    "sync"
)

// 结果：fatal error: all goroutines are asleep - deadlock!

// 重入锁：当一个线程获取锁时，如果没有其它线程拥有这个锁，那么，这个线程就成功获取到这个锁。
// 之后，如果其它线程再请求这个锁，就会处于阻塞等待的状态。
// 但是，如果拥有这把锁的线程再请求这把锁的话，不会阻塞，而是成功返回，所以叫可重入锁(有时候也叫做递归锁)。

// 解析：
func foo(l sync.Locker) {
    fmt.Println("in foo")
    l.Lock()
    bar(l)
    l.Unlock()
}
func bar(l sync.Locker) {
    l.Lock()
    fmt.Println("in bar")
    l.Unlock()
}
func main() {
    l := &sync.Mutex{}
    foo(l)
}
