package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var (
		t1 = "hello world"
		t2 = `"hello World"`
		t3 = "\"hello World\""
	)

	a := 3
	var c int = 3
	var s string = "湖心小筑"
	const b = 2
	//s=string(b)
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(s)
	fmt.Printf(s)
	fmt.Printf("\n")
	fmt.Printf(t1)
	fmt.Printf("\n")
	fmt.Printf(t2)
	fmt.Printf("\n")
	fmt.Printf(t3)
	fmt.Printf("\n")
	fmt.Print(c)
	fmt.Println("------------")
	fmt.Print(t1,t2,t3)
}
