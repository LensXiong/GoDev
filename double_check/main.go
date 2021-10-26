package double_check

import (
    "sync"
)

// 在多核 CPU 中，因为 CPU 缓存会导致多个核心中变量值不同步。
type Once struct {
    m    sync.Mutex
    done uint32
}

func (o *Once) Do(f func()) {
    if o.done == 1 {
        return
    }
    o.m.Lock()
    defer o.m.Unlock()
    if o.done == 0 {
        o.done = 1
        f()
    }
}
