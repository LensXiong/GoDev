package main

// 要点：基于 Mutex 实现 TryLock。

// 当一个 goroutine 调用这个 TryLock 方法请求锁的时候，如果这把锁没有被其他 goroutine 所持有.
// 那么，这个 goroutine 就持有了这把锁，并返回 true;
// 如果这把锁已经被其他 goroutine 所持有，或者是正在准备交给某个被唤醒的 goroutine。
// 那么这个请求锁的 goroutine 就直接返回 false，不会阻塞在方法调用上。
import (
    "fmt"
    "math/rand"
    "sync"
    "sync/atomic"
    "time"
    "unsafe"
)

// 复制Mutex定义的常量
const (
    mutexLocked      = 1 << iota // 加锁标识位置
    mutexWoken                   // 唤醒标识位置
    mutexStarving                // 锁饥饿标识位置
    mutexWaiterShift = iota      // 标识waiter的起始bit位置
)

// 扩展一个Mutex结构
type Mutex struct {
    sync.Mutex
}

// 尝试获取锁
func (m *Mutex) TryLock() bool {
    // 如果能成功抢到锁
    if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexWaiterShift) {
        return true
    }
    // 如果处于唤醒、加锁或者饥饿状态，这次请求就不参与竞争了，返回false
    old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
    if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
        return false
    }
    // 尝试在竞争的状态下请求锁
    new := old | mutexLocked
    return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new)
}

func main() {
    var mu Mutex
    go func() { // 启动一个goroutine持有一段时间的锁
        mu.Lock()
        time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
        mu.Unlock()
    }()
    time.Sleep(time.Second)
    ok := mu.TryLock() // 尝试获取到锁
    if ok {            //获取成功
        fmt.Println("got the lock")
        // do something
        mu.Unlock()
        return
    }

    // 没有获取到
    fmt.Println("can't get the lock")
}
