package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

const (
	a = "this is a"
	b = "this is b"
)

func main() {
	fmt.Printf("%b\t", kb)
	fmt.Printf("%d\n", kb)
	fmt.Printf("%b\t", mb)
	fmt.Printf("%d\n", mb)
	fmt.Printf("%b\t", gb)
	fmt.Printf("%d\n", gb)

	fmt.Println(a)
	fmt.Println(b)
}
