package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
    "unsafe"
)

// 要点：获取等待者的数量等指标。
// 通过 unsafe 的方式读取到 Mutex 内部的 state 字段。

const (
    mutexLocked = 1 << iota // mutex is locked
    mutexWoken
    mutexStarving
    mutexWaiterShift = iota // iota=3
)

type Mutex struct {
    sync.Mutex
}

func (m *Mutex) Count() int {
    // 获取state字段的值
    v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
    v = v >> mutexWaiterShift // 得到等待者的数值
    v = v + (v & mutexLocked) // 再加上锁持有者的数量，0或者1

    return int(v)
}

// 锁是否被持有
func (m *Mutex) IsLocked() bool {
    state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
    return state&mutexLocked == mutexLocked

}

// 是否有等待者被唤醒
func (m *Mutex) IsWoken() bool {
    state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
    return state&mutexWoken == mutexWoken
}

// 锁是否处于饥饿状态
func (m *Mutex) IsStarving() bool {
    state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
    return state&mutexStarving == mutexStarving
}

func main() {
    var mu Mutex
    for i := 0; i < 1000; i++ { // 启动1000个goroutine
        go func() {
            mu.Lock()
            time.Sleep(time.Nanosecond)
            mu.Unlock()
        }()

        time.Sleep(time.Nanosecond)
        // 输出锁的信息
        fmt.Printf("waitings: %d, isLocked: %t, woken: %t, starving: %t\n",
            mu.Count(), mu.IsLocked(), mu.IsWoken(), mu.IsStarving())
    }
}
