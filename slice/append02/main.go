package main

import (
"fmt"
)

// 结果：[0 0 0 0 0 0 0 0 0 0 1 2 3]
// 解析：make 初始化均为0；
// append 操作的本质就是对数组扩容；
// go 底层会创建一个新的数组 newArr(按照扩容后大小) 将 slice原来包含的元素拷贝到新的数组 newArr， slice 重新引用到 newArr。
func main() {
    s := make([]int, 10)

    s = append(s, 1, 2, 3)

    fmt.Println(s)
}
