package main

import "fmt"

/*
   下面代码是否编译通过?
*/

// 结果：syntax error: mixed named and unnamed function parameters

// 解析：在函数有多个返回值时，只要有一个返回值有指定命名，其他的也必须有命名。
// 如果返回值有有多个返回值必须加上括号；
// 如果只有一个返回值并且有命名也需要加上括号；
// 此处函数第一个返回值有 sum 名称，第二个未命名，所以错误。

//func myFunc(x, y int) (sum int, error) {
//    return x + y, nil
//}

func myFunc(x, y int) (sum int, err error) {
    return x + y, nil
}
func main() {
    num, _ := myFunc(1, 2)
    fmt.Println("num = ", num)
}
