package main

import "fmt"

// 知识点：关于 map 的遍历赋值，for range 中，stu 是结构体的一个拷贝副本，
// 所以 m[stu.Name]=&stu 实际上一致指向同一个指针， 最终该指针的值为遍历的最后一个 struct 的值拷贝。

// 结果：
// li => wang
// wang => wang
// zhou => wang

// 逃逸分析 go build -gcflags '-m -l' ./main.go
// # command-line-arguments
// ./main.go:27:12: moved to heap: stu
// ./main.go:17:14: make(map[string]*student) does not escape
// ./main.go:20:26: []student literal does not escape
// ./main.go:38:20: ... argument does not escape
// ./main.go:38:20: k escapes to heap
// ./main.go:38:24: "=>" escapes to heap
// ./main.go:38:31: v.Name escapes to heap

type Student struct {
    Name string
    Age  int
}

func main() {
    // 定义map
    m := make(map[string]*Student)

    // 定义student数组
    students := []Student{
        {Name: "zhou", Age: 24},
        {Name: "li", Age: 23},
        {Name: "wang", Age: 22},
    }

    // 将数组依次添加到map中（错误写法）
    for _, stu := range students {
        m[stu.Name] = &stu
    }

    // 遍历结构体数组，依次赋值给map （正确写法）
    // for i := 0; i < len(students); i++  {
    //    m[students[i].Name] = &students[i]
    // }

    // 打印map
    for k, v := range m {
        fmt.Println(k, "=>", v.Name)
    }
}
