package main

// 逃逸场景二：向 channel 发送指针数据。
// 因为在编译时，不知道 channel 中的数据会被哪个 goroutine 接收，因此编译器没法知道变量什么时候才会被释放，因此只能放入堆中。

// 结果：go build -gcflags '-m -l' ./main.go
// # command-line-arguments
// ./main.go:12:5: moved to heap: y
func main() {
    ch := make(chan int, 1)
    x := 5
    ch <- x // x 不发生逃逸，因为只是复制的值
    ch1 := make(chan *int, 1)
    y := 5
    py := &y
    ch1 <- py // y 逃逸，因为 y 地址传入了 chan 中，编译时无法确定什么时候会被接收，所以也无法在函数返回后回收y
}
