package main

// 结果：go build -gcflags '-m -l' ./main.go
// moved to heap: fooVal3

// 0xc00002e758 0xc00002e738 0xc00002e730 0xc00008e000 0xc00002e728 0xc00002e720
// 0xc00002e758 0xc00002e738 0xc00002e730 0xc00008e000 0xc00002e728 0xc00002e720
// 0xc00002e758 0xc00002e738 0xc00002e730 0xc00008e000 0xc00002e728 0xc00002e720
// 0xc00002e758 0xc00002e738 0xc00002e730 0xc00008e000 0xc00002e728 0xc00002e720
// 0xc00002e758 0xc00002e738 0xc00002e730 0xc00008e000 0xc00002e728 0xc00002e720
// 13 0xc00008e000

func foo(argVal int) *int {

    var fooVal1 int = 11
    var fooVal2 int = 12
    var fooVal3 int = 13
    var fooVal4 int = 14
    var fooVal5 int = 15

    // 此处循环是防止go编译器将foo优化成inline(内联函数)
    // 如果是内联函数，main调用foo将是原地展开，所以fooVal1-5相当于main作用域的变量
    // 即使fooVal3发生逃逸，地址与其他也是连续的
    for i := 0; i < 5; i++ {
        println(&argVal, &fooVal1, &fooVal2, &fooVal3, &fooVal4, &fooVal5)
    }

    // 返回fooVal3给main函数
    return &fooVal3
}

func main() {
    mainVal := foo(666)

    println(*mainVal, mainVal)
}
