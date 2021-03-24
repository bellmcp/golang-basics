package main

import "fmt"

// 0xc00010c008 (memory address)		20
func main() {
	// var pointer *int // *int = pointer at int
	// i := 20

	// pointer = &i // get the memory address of i (0xc00010c008)
	// fmt.Println(pointer)
	// fmt.Println(*pointer) // get value from the memory address of i (20)

	i := 20
	pointer := &i // SHORT HAND get the pointer at int i
	fmt.Println(pointer)
	fmt.Println(*pointer)
}
