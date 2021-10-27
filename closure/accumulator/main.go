package main

import "fmt"

// 结果：
// (0xc00001c0a0, 0) - 1
// (0xc00001c0a0, 1) - 11
// (0xc00001c0a0, 11) - 111
// ------------------------
// (0xc00001c0c0, 0) - 1
// (0xc00001c0c0, 1) - 11
// (0xc00001c0c0, 11) - 111

// 解析：闭包引用了x变量，a,b可看作2个不同的实例，实例之间互不影响。实例内部，x变量是同一个地址，因此具有累加效应。
// 拓展：闭包是由函数及其相关引用环境组合而成的实体，即：闭包 = 函数 + 引用环境
// 一般的函数都有函数名，但是匿名函数就没有。匿名函数不能独立存在，但可以直接调用或者赋值于某个变量。
// 匿名函数也被称为闭包，一个闭包继承了函数声明时的作用域。在 Golang 中，所有的匿名函数都是闭包。
// 可以把闭包看成是一个类，一个闭包函数调用就是实例化一个类。
// 闭包在运行时可以有多个实例，它会将同一个作用域里的变量和常量捕获下来，无论闭包在什么地方被调用（实例化）时，都可以使用这些变量和常量。
// 而且，闭包捕获的变量和常量是引用传递，不是值传递。

func main() {
    var a = Accumulator()
    fmt.Printf("%d\n", a(1))
    fmt.Printf("%d\n", a(10))
    fmt.Printf("%d\n", a(100))
    fmt.Println("---------------")

    var b = Accumulator()
    fmt.Printf("%d\n", b(1))
    fmt.Printf("%d\n", b(10))
    fmt.Printf("%d\n", b(100))
}

func Accumulator() func(int) int {
    var x int
    return func(delta int) int {
        fmt.Printf("(%+v, %+v) - ", &x, x)
        x += delta
        return x
    }
}
