package main

import "sync"

// 问题：如何实现线程安全的 map 类型?

// 三种方案：
// 方案①：加读写锁，扩展 map，支持并发读写。
// 方案②：分片加锁（更高效的并发`map`）。
// 方案③：应对特殊场景的 sync.Map。

// 方案一解析：避免 map 并发读写 panic 的方式之一就是加锁，考虑到读写性能，可以使用读写锁提供性能。
// 对 map 对象的操作，无非就是增删改查和遍历等几种常见操作。
// 可以把这些操作分为读和写两类，其中，查询和遍历可以看做读操作，增加、修改和删除可以看做写操作。
// 通过读写锁对相应的操作进行保护。

// 读写锁的缺点。虽然使用读写锁可以提供线程安全的 map，但是在大量并发读写的情况下，锁的竞争会非
// 常激烈，毕竟锁是性能下降的万恶之源之一。因此，需要尽量减少锁的粒度和锁的持有时间。

type RWMap struct { // 一个读写锁保护的线程安全的map
    sync.RWMutex // 读写锁保护下面的map字段
    m            map[int]int
}

func (m *RWMap) Get(k int) (int, bool) { //从map中读取一个值
    m.RLock()
    defer m.RUnlock()
    v, existed := m.m[k] // 在锁的保护下从map中读取
    return v, existed
}

func (m *RWMap) Set(k int, v int) { // 设置一个键值对
    m.Lock() // 锁保护
    defer m.Unlock()
    m.m[k] = v
}

func (m *RWMap) Delete(k int) { //删除一个键
    m.Lock() // 锁保护
    defer m.Unlock()
    delete(m.m, k)
}

func (m *RWMap) Len() int { // map的长度
    m.RLock() // 锁保护
    defer m.RUnlock()
    return len(m.m)
}

func (m *RWMap) Each(f func(k, v int) bool) { // 遍历map
    m.RLock() //遍历期间一直持有读锁
    defer m.RUnlock()
    for k, v := range m.m {
        if !f(k, v) {
            return
        }
    }
}
