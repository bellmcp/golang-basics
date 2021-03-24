package main

import "fmt"

// func inc(num int) int {
// 	return num + 1
// }

// func inc(num int) int {
// 	return num + 1
// }

func inc(num *int) {
	*num++ //*num = *num + 1
	// -> *num = get the value of the address of num
}

func main() {
	i := 20
	// result := inc(i) // send the copied of i to the function  -> PASSED BY VALUE

	inc(&i) // send the copied of memory address of i to the function -> PASSED BY VALUE

	fmt.Println(i) //21
}
