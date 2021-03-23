package main

import "fmt"

func main() {
	// const pi = 3.14 //immutable
	// ** GoLang convention does not use uppercase const like other languages

	// MULTIPLE DECLARATION
	// const (
	// 	i = 1
	// 	j // j will have the same value as i
	// )

	// const (
	// 	red   = iota // 0 -> automatic order declaration start with 0
	// 	green        // 1
	// 	blue         // 2
	// )

	// const (
	// 	red   = iota + 1 // 1
	// 	green            // 2
	// 	blue             // 3
	// )

	const (
		sun = iota + 1 // 1
		mon            // 2
		tue            // 3
		_              // skip
		_              // skip
		_              // skip
		sat            // 7
	)

	fmt.Println(sat)
}
