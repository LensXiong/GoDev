package main

const cl = 100

var bl = 123

// 结果：cannot take the address of cl
// 解析：常量是无法取出地址的，因为字面量符号并没有地址而言。
// 常量不同于变量的在运行期分配内存，常量通常会被编译器在预处理阶段直接展开，作为指令数据使用。
func main() {
    println(&bl, bl)
    println(&cl, cl)
}
