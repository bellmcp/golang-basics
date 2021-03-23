package main

import "fmt"

func main() {
	// for n := 1; n <= 10; n++ { // declare; condition; postcondition
	// 	fmt.Println(n)
	// }

	// There is no while in GoLang -> use for
	n := 1        // declare
	for n <= 10 { // _ ; condition; _
		fmt.Println(n)
		n++ // postcondition
	}

	// // INFINITE FOR (FOREVER)
	// for {
	// 	fmt.Println("Hello")
	// }
}
