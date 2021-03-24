package main

import "fmt"

// func add(a int, b int) int { // outer int = type of returned value
// 	return a + b
// }

func add(a, b int) int { // grouped inner int
	return a + b
}

func next(num int) (int, int) { // return more than one value
	return num + 1, num + 2
}

func main() {
	// result := add(10, 20)
	// fmt.Println(result)

	x, y := next(10) // x = 11, y =12
	fmt.Println(x, y)
}
