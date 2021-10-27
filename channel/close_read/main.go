package main

import "fmt"

// 问题：对已经关闭的的 chan 进行读写，会怎么样?为什么?
// ① 写已经关闭的 chan 会 panic。报错信息：panic: send on closed channel。
// ② 读已经关闭的 chan 能一直读到东⻄，但是读到的内容根据通道内关闭前是否有元素而不同。
// 如果 chan 关闭前，buffer 内有元素还未读 , 会正确读到 chan 内的值，且返回的第二个 bool 值(是否读成功)为 true。
// 如果 chan 关闭前，buffer 内有元素已经被读完，chan 内无值，接下来所有接收的值都会非阻塞直接成功，返回 channel 元素的零值，第二个 bool 值一直为 false。

// 结果：
// 以下是数值的chan
// 第一次读chan的协程结束，num=1， ok=true
// 第二次读chan的协程结束，num=0， ok=false
// 第三次读chan的协程结束，num=0， ok=false

// 以下是字符串chan
// 第一次读chan的协程结束，str=aaa， ok=true
// 第二次读chan的协程结束，str=， ok=false
// 第三次读chan的协程结束，str=， ok=false

// 以下是结构体chan
// 第一次读chan的协程结束，struct={ha}， ok=true
// 第二次读chan的协程结束，struct={}， ok=false
// 第三次读chan的协程结束，struct={}， ok=false

func main() {
    fmt.Println("以下是数值的chan")
    ci := make(chan int, 3)
    ci <- 1
    close(ci)
    num, ok := <-ci
    fmt.Printf("第一次读chan的协程结束，num=%v， ok=%v\n", num, ok)
    num1, ok1 := <-ci
    fmt.Printf("第二次读chan的协程结束，num=%v， ok=%v\n", num1, ok1)
    num2, ok2 := <-ci
    fmt.Printf("第三次读chan的协程结束，num=%v， ok=%v\n", num2, ok2)
    fmt.Println()

    fmt.Println("以下是字符串chan")
    cs := make(chan string, 3)
    cs <- "aaa"
    close(cs)
    str, ok := <-cs
    fmt.Printf("第一次读chan的协程结束，str=%v， ok=%v\n", str, ok)
    str1, ok1 := <-cs
    fmt.Printf("第二次读chan的协程结束，str=%v， ok=%v\n", str1, ok1)
    str2, ok2 := <-cs
    fmt.Printf("第三次读chan的协程结束，str=%v， ok=%v\n", str2, ok2)
    fmt.Println()

    fmt.Println("以下是结构体chan")
    type MyStruct struct {
        Name string
    }
    cst := make(chan MyStruct, 3)
    cst <- MyStruct{Name: "ha"}
    close(cst)
    struct1, ok := <-cst
    fmt.Printf("第一次读chan的协程结束，struct=%v， ok=%v\n", struct1, ok)
    struct2, ok1 := <-cst
    fmt.Printf("第二次读chan的协程结束，struct=%v， ok=%v\n", struct2, ok1)
    struct3, ok2 := <-cst
    fmt.Printf("第三次读chan的协程结束，struct=%v， ok=%v\n", struct3, ok2)
}
