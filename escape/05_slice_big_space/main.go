package main

// 逃逸场景五：切片（扩容后）长度太大
// 解析：实际上当栈空间不足以存放当前对象时或无法判断当前切片长度时会将对象分配到堆中。
// 结果： go build -gcflags '-m -l' ./main.go
// # command-line-arguments
// ./main.go:8:14: make([]int, 10000, 10000) escapes to heap

func main() {
    Slice() // 会发生逃逸吗？
}

func Slice() {
    s := make([]int, 10000, 10000)

    for index, _ := range s {
        s[index] = index
    }
}
