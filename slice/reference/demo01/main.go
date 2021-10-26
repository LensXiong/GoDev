package main

import (
    "fmt"
)

// 解析：golang 切片是引用类型，所以在传递时，遵守引用传递机制。
// 当使用 str1[1:] 时，str2 和 str1 底层共享一个数组，这回导致 str2[1] = "new" 语句影响 str1 。

// 切片 append 操作的底层原理分析: append 操作的本质就是对数组扩容
// go 底层会创建一下新的数组 newArr(按照扩容后大小) 将 slice 原来包含的元素拷贝到新的数组 newArr ，原来的 slice 重新引用到 newArr
// append 会导致底层数组扩容，生成新的数组，因此追加数据后的 str2 不会影响 str1 。
func main() {
    str1 := []string{"a", "b", "c"}
    str2 := str1[1:]
    str2[1] = "new"
    fmt.Println(str1) // [a b new]
    fmt.Println(str2) // [b new]
    str2 = append(str2, "z", "x", "y")
    fmt.Println(str1) // [a b new]
    fmt.Println(str2) // [b new z x y]
}
