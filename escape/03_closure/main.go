package main

// 逃逸场景三：局部变量在函数调用结束后还被其他地方（闭包中引用包外的值或者函数返回局部变量指针）使用。
// 因为变量的生命周期可能会超过函数周期，因此只能放入堆中。

// 结果：# command-line-arguments
// ./main.go:7:5: moved to heap: x
// ./main.go:8:12: func literal escapes to heap
func Foo() func() {
    x := 5 // x 发生逃逸，因为在 Foo 调用完成后，被闭包函数用到，还不能回收，只能放到堆上存放
    return func() {
        x += 1
    }
}
func main() {
    inner := Foo()
    inner()
}
