package main

// 知识点：Channel happens-before 关系保证有 4 条规则。

// 知识点：对于 unbuffered 的 Channel，也就是容量是 0 的 Channel，
// 从此 Channel 中读取数据的调用一定 happens before 往此 Channel 发送数据的调用完成。

// 结果：hello,world
var ch = make(chan int)

var s string

func f() {
    s = "hello,world"
    <-ch
}

func main() {
    go f()
    ch <- 1
    print(s)
}
