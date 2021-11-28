package main

import "fmt"

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

type student struct {
    Name string
    Age  int
}

func main() {
    // 定义map
    m := make(map[string]*student)

    // 定义student数组
    students := []student{
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
