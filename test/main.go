package main

import "fmt"

//func main() {
//    arr := make([]int, 0)
//    for i := 0; i < 2000; i++ {
//        fmt.Println("len 为", len(arr), "cap 为", cap(arr))
//        arr = append(arr, i)
//    }
//
//}

//// 1 1 2 3 5 8 13 21 34 55
//func main (){
//    slice := make([]int,10)
//    for i := 0; i< 10; i++{
//        if i < 2{
//            slice[i] =  1
//        } else {
//            slice[i] = slice[i-2] + slice[i-1]
//        }
//
//    }
//    fmt.Println("slice ",slice)
//}


/*
[{
	"name": "aaa",
	"children": [{
		"name": "bbb",
		"children": []
	}]
}, {
	"name": "ccc",
	"children": []
}]

[{
	"name": "aaa",
	"children": []
}, {
	"name": "bbb",
	"children": []
}, {
	"name": "ccc",
	"children": []
}]

 */


type T struct {
	a int
	b float64
	c string
}
func main (){
	t := &T{ 7, -2.35, "abc\tdef" }
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)

	x := []int{1,2,3}
	y := []int{4,5,6}
	x = append(x, y...)
	fmt.Println(x)
}



