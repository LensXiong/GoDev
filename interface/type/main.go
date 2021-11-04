package main

type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

// cannot use s (type S) as type *interface {} in argument to g:
//	*interface {} is pointer to interface, not interface

// cannot use p (type *S) as type *interface {} in argument to g:
//	*interface {} is pointer to interface, not interface
func main() {
    s := S{}
    p := &s
    f(s) //A
    g(s) //B
    f(p) //C
    g(p) //D
}
