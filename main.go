package main

import (
	"fmt"
	"runtime"
)

func main() {
	// num := 21

	// if num%2 == 0 {
	// 	fmt.Println("Even")
	// } else {
	// 	fmt.Println("Odd")
	// }

	// os := runtime.GOOS

	// if os == "darwin" {
	// 	fmt.Println("You cool")
	// } else if os == "windows" {
	// 	fmt.Println("Switch to mac")
	// } else {
	// 	fmt.Println("T_T")
	// }

	// MERGE DECLARATION AND CONDITION
	if os := runtime.GOOS; os == "darwin" { // use ;
		fmt.Println("You cool")
	} else if os == "windows" {
		fmt.Println("Switch to mac")
	} else {
		fmt.Println("T_T")
	}
}
