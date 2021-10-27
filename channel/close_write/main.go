package main

// 问题：对已经关闭的的 chan 进行读写，会怎么样?为什么?
// ① 写已经关闭的 chan 会 panic。报错信息：panic: send on closed channel。
// ② 读已经关闭的 chan 能一直读到东⻄，但是读到的内容根据通道内关闭前是否有元素而不同。
// 如果 chan 关闭前，buffer 内有元素还未读 , 会正确读到 chan 内的值，且返回的第二个 bool 值(是否读成功)为 true。
// 如果 chan 关闭前，buffer 内有元素已经被读完，chan 内无值，接下来所有接收的值都会非阻塞直接成功，返回 channel 元素的零值，但是第二个 bool 值一直为 false。

// 追加：为什么写已经关闭的 chan 就会 panic 呢?
// 解析: 在 src/runtime/chan.go源码
/**
func chansend(c *hchan,ep unsafe.Pointer,block bool,callerpc uintptr) bool {
    ...
    if c.closed != 0 {
        unlock(&c.lock)
        panic(plainError("send on closed channel"))
}
    ...
}
*/

func main() {
    c := make(chan int, 3)
    close(c)
    c <- 1
}
