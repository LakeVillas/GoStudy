package main

import (
	"fmt"
	"strconv"
)

var (
	s1  string
	s2  string
	b   int
	c   int
	arr [4]int
)

func main() {
	a := 120
	b = a << 2
	c = a >> 3
	s1 = "hello"
	s2 = "world"
	//s4=s1+s2
	//s3={1,3,5}
	//fmt.Print(s3)
	arr[3] = 10
	fmt.Print(s1[0])
	fmt.Println(s2)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%d 的二进制是 %b", a, a)
	fmt.Println()
	//s5:=strconv.FormatInt(int64(a),2)
	s6 := strconv.FormatInt(int64(111100000), 10)
	fmt.Println("Hello" + "World")
	fmt.Println("Hello"[3])
	fmt.Println(s6)
	fmt.Println(arr[3])
	fmt.Println(len(arr))
	f1()
}

func f1()  {
	for i:=0;i<100;i++ {
		fmt.Println(i)
	}
}