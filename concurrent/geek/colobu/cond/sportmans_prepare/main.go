package main

import (
    "log"
    "math/rand"
    "sync"
    "time"
)

//  Go 标准库提供 Cond 原语的目的是，为等待 / 通知场景下的并发问题提供支持。
//  Cond 通常应用于等待某个条件的一组 goroutine，等条件变为 true 的时候，其中一个 goroutine 或者所有的 goroutine 都会被唤醒执行。

// Cond 是和某个条件相关，这个条件需要一组 goroutine 协作共同完成。
// 在条件还没有满足的时候，所有等待这个条件的 goroutine 都会被阻塞住，
// 只有这一组 goroutine 通过协作达到了这个条件，等待的 goroutine 才可能继续进行下去。

// 使用 Cond 的 2 个常见错误：
// ① 调用 Wait 的时候没有加锁。
// ② 只调用了一次 Wait，没有检查等待条件是否满足，结果条件没满足，程序就继续执行了。

// 错误①分析：在调用 cond.Wait 时，如果把前后的 Lock/Unlock 注释掉，再运行程序，
// 就会报释放未加锁的 panic:fatal error: sync: unlock of unlocked mutex
// 原因：cond.Wait 方法的实现是，把当前调用者加入到 notify 队列之中后会释放锁(如果不释放锁，
// 其他 Wait 的调用者就没有机会加入到 notify 队列中 了)，然后一直等待;等调用者被唤醒之后，又会去争抢这把锁。
// 如果调用 Wait 之前不加锁的话，就有可能 Unlock 一个未加锁的 Locker。所以切记，调用 cond.Wait 方法之前一定要加锁。

// 错误②分析：将 ready 条件去掉。
// 运行这个程序，你会发现，可能只有几个运动员准备好之后程序就运行完了，而不是我们期望的所有运动员都准备好才进行下一步。
// 原因在于，每一个运动员准备好之后都会唤醒所有的等待者，也就是这里的裁判员，比如第一个运动员准备好后就唤醒了裁判员，结果
// 这个裁判员傻傻地没做任何检查，以为所有的运动员都准备好了，就继续执行了。

// waiter goroutine 被唤醒不等于等待条件被满足，只是有 goroutine 把它唤醒了而已，
// 等待条件有可能已经满足了，也有可能不满足，我们需要进一步检查。等待者被唤醒，只是得到了一次检查的机会而已。
func main() {
    c := sync.NewCond(&sync.Mutex{})
    var ready int

    for i := 1; i < 11; i++ {
        go func(i int) {
            time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

            // 加锁更改等待条件
            c.L.Lock()
            ready++
            c.L.Unlock()

            log.Printf("运动员#%d 已准备就绪\n", i)
            // 广播唤醒所有的等待者
            c.Broadcast()
        }(i)
    }

    c.L.Lock()
    for ready != 10 {
        c.Wait()
        log.Println("裁判员被唤醒一次")
    }
    c.L.Unlock()

    //所有的运动员是否就绪
    log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")
}
