package main

import "fmt"

type mapKey struct {
    key int
}

// 知识点：map中key的类型选择。
// map 的类型是 map[key]，key 类型的 K 必须是可比较的，通常情况下，我们会选择内建的基本类型，比如整数、字符串做 key 的 类型。
// 如果要使用 struct 作为 key，我们要保证 struct 对象在逻辑上是不可变的。

// 运行结果：key={10}
// m[key]=hello
// 再次查询m[key]=

// 解析：如果使用 struct 类型做 key 其实是有坑的，因为如果 struct 的某个字段值修改了，查询 map 时无法获取它 add 进去的值。

func main() {
    var m = make(map[mapKey]string)
    var key = mapKey{10}

    fmt.Printf("key=%d\n", key)

    m[key] = "hello"
    fmt.Printf("m[key]=%s\n", m[key])

    // 修改key的字段的值后再次查询map，无法获取刚才add进去的值
    key.key = 100
    fmt.Printf("再次查询m[key]=%s\n", m[key])
}
