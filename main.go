package main

import "fmt"

func main() {
	words := []string{"Hello", "Hi", "Bye", "Wutipat", "Bell"}

	// FIND
	fmt.Println(words[2])

	// APPEND
	words = append(words, "Go")
	fmt.Println(words)

	// REMOVE -> "Bye"
	// {"Hello", "Hi"} + {"Wutipat", "Bell"}
	// words = append(words[:2], words[3:]) //start with index 0 to 1 (except words[2]) then merges with the rest
	// ** CANNOT DO THIS -> it cannot concat 2 slices

	// DO THIS
	words = append(words[:2], words[3:]...) // ... behind slices is like spread operator in JS
	fmt.Println(words)
}
