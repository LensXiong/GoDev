package main

// Cond 有三个方法分别是 Broadcast、Signal 和 Wait 方法。

// Signal 方法，允许调用者 Caller 唤醒一个等待此 Cond 的 goroutine。如果此时没有等待的 goroutine，显然无需通知 waiter;
// 如果 Cond 等待队列中有一个或者多个等待的 goroutine，则需要从等待队列中移除第一个 goroutine 并把它唤醒。
// 调用 Signal 方法时，不强求你一定要持有 c.L 的锁。

// Broadcast 方法，允许调用者 Caller 唤醒所有等待此 Cond 的 goroutine。如果此时没有等待的 goroutine，显然无需通知 waiter;
// 如果 Cond 等待队列中有一个或者多个等待的 goroutine，则清空所有等待的 goroutine，并全部唤醒。
// 同样地，调用 Broadcast 方法时，也不强求一定持有 c.L 的锁。

// Wait 方法，会把调用者 Caller 放入 Cond 的等待队列中并阻塞，直到被 Signal 或者 Broadcast 的方法从等待队列中移除并唤醒。
// 调用 Wait 方法时必须要持有 c.L 的锁。

/**
type Cond struct {
    noCopy noCopy

    // 当观察或者修改等待条件的时候需要加锁
    L Locker

    // 等待队列
    notify  notifyList
    checker copyChecker
}

func NewCond(l Locker) *Cond {

    return &Cond{L: l}

}

func (c *Cond) Wait() {

    c.checker.check()
    // 增加到等待队列中

    t := runtime_notifyListAdd(&c.notify)

    c.L.Unlock()
    // 阻塞休眠直到被唤醒

    runtime_notifyListWait(&c.notify, t)

    c.L.Lock()

}

func (c *Cond) Signal() {

    c.checker.check()

    runtime_notifyListNotifyOne(&c.notify)

}

func (c *Cond) Broadcast() {

    c.checker.check()

    runtime_notifyListNotifyAll(&c.notify)

}
**/
