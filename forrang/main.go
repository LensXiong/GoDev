package main

import "fmt"

type student struct {
    Name string
    Age  int
}

func iterateStudent() {
    m := make(map[string]*student)
    s := []student{
        {"gao", 24},
        {"li", 23},
        {"wang", 22},
    }
    for _, stu := range s {
        fmt.Printf("%v", stu)
        fmt.Println()
        m[stu.Name] = &stu
    }
    fmt.Println()
    fmt.Printf("%v", m)
}

func main() {
    iterateStudent()
}
