package main

import (
    "fmt"
    "sync"
    "time"
)

// Mutex 常见的问题之4-死锁问题。

// 运行结果：fatal error: all goroutines are asleep - deadlock!

// 保证 Lock/Unlock 成对出现，尽可能采用 defer mutex.Unlock 的方式，把它们成对、紧凑地写在一起。

// 解析：两个或两个以上的进程(或线程，goroutine)在执行过程中，因争夺共享资源而处于一种互相等待的状态，
// 如果没有外部干涉，它们都将无法推进下去， 此时，我们称系统处于死锁状态或系统产生了死锁。

// 死锁产生的必要条件：
// ① 互斥: 至少一个资源是被排他性独享的，其他线程必须处于等待状态，直到资源被释放。
// ② 持有和等待: goroutine 持有一个资源，并且还在请求其它 goroutine 持有的资源，也就是咱们常说的“吃着碗里，看着锅里”的意思。
// ③ 不可剥夺: 资源只能由持有它的 goroutine 来释放。
// ④ 环路等待: 一般来说，存在一组等待进程，P={P1，P2，...，PN}，P1 等待 P2 持有的 资源，P2 等待 P3 持有的资源，
// 依此类推，最后是 PN 等待 P1 持有的资源，这就形成了一个环路等待的死结。

func main() {
    // 派出所证明
    var psCertificate sync.Mutex
    // 物业证明
    var propertyCertificate sync.Mutex
    var wg sync.WaitGroup
    wg.Add(2) // 需要派出所和物业都处理
    // 派出所处理 goroutine1
    go func() {
        defer wg.Done() // 派出所处理完成
        psCertificate.Lock()
        defer psCertificate.Unlock()
        // 检查材料
        time.Sleep(5 * time.Second) // 请求物业的证明
        propertyCertificate.Lock()
        propertyCertificate.Unlock()
    }()
    // 物业处理 goroutine2
    go func() {
        defer wg.Done() // 物业处理完成
        propertyCertificate.Lock()
        defer propertyCertificate.Unlock()
        // 检查材料
        time.Sleep(5 * time.Second) // 请求派出所的证明
        psCertificate.Lock()
        psCertificate.Unlock()
    }()
    wg.Wait()
    fmt.Println("成功完成")
}
