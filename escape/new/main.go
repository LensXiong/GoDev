package main

// 结果：go build -gcflags '-m -l' ./main.go
// moved to heap: fooVal3

// 666 0xc00002e728 0xc00002e720 0xc000110000 0xc00002e738 0xc00002e730
// 666 0xc00002e728 0xc00002e720 0xc000110000 0xc00002e738 0xc00002e730
// 666 0xc00002e728 0xc00002e720 0xc000110000 0xc00002e738 0xc00002e730
// 666 0xc00002e728 0xc00002e720 0xc000110000 0xc00002e738 0xc00002e730
// 666 0xc00002e728 0xc00002e720 0xc000110000 0xc00002e738 0xc00002e730
// 0 0xc000110000

func foo(argVal int) *int {

    var fooVal1 *int = new(int)
    var fooVal2 *int = new(int)
    var fooVal3 *int = new(int)
    var fooVal4 *int = new(int)
    var fooVal5 *int = new(int)

    // 此处循环是防止go编译器将foo优化成inline(内联函数)
    // 如果是内联函数，main调用foo将是原地展开，所以fooVal1-5相当于main作用域的变量
    // 即使fooVal3发生逃逸，地址与其他也是连续的
    for i := 0; i < 5; i++ {
        println(argVal, fooVal1, fooVal2, fooVal3, fooVal4, fooVal5)
    }

    // 返回fooVal3给main函数
    return fooVal3
}

func main() {
    mainVal := foo(666)

    println(*mainVal, mainVal)
}
