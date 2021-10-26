package main

import (
    "sync"
    "time"
)

// 解析: WaitGroup 在调用 Wait 之后是不能再调用 Add 方法的。
// 如果 Wait 方法执行期间，跨越了两个计数周期，就会引发 panic 。
// 例如，当前 goroutine 调用 Wait 方法阻塞了。另外一个 goroutine 调用了 Done 方法，计数器值变为0；
// 此时会唤醒当前 goroutine，并且试图继续执行 Wait 方法中其余代码。
// 这时，又有一个 goroutine 调用了 Add 方法，此时 Wait 方法就会抛出：
// panic: sync: WaitGroup is reused before previous Wait has returned

// Add adds delta, which may be negative, to the WaitGroup counter.
// If the counter becomes zero, all goroutines blocked on Wait are released.
// If the counter goes negative, Add panics.

// Note that calls with a positive delta that occur when the counter is zero
// must happen before a Wait. Calls with a negative delta, or calls with a
// positive delta that start when the counter is greater than zero, may happen
// at any time.
// Typically this means the calls to Add should execute before the statement
// creating the goroutine or other event to be waited for.
// If a WaitGroup is reused to wait for several independent sets of events,
// new Add calls must happen after all previous Wait calls have returned.
// See the WaitGroup example.

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        time.Sleep(time.Millisecond)
        wg.Done()
        wg.Add(1)
    }()
    wg.Wait()
}
