package main

// 逃逸场景四：在 slice 或 map 中存储指针。
// 比如 []*int，其后面的数组可能是在栈上分配的，但其引用的值还是在堆上。

// 结果： go build -gcflags '-m -l' ./main.go
// # command-line-arguments
// ./main.go:6:9: moved to heap: x

func main() {
    var x int
    x = 10
    var ls []*int
    ls = append(ls, &x) // x发生逃逸，ls存储的是指针，所以ls底层的数组虽然在栈存储，但x本身却是逃逸到堆上
}
