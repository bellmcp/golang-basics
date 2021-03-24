package main

import "fmt"

func main() {
	// arr := [3]string{"A", "B", "C"}
	// arr := [...]string{"A", "B", "C"} -> ... = auto assign array size
	// arr[1] = "B"

	// ASSIGN SOME VALUE
	// arr := [3]string{1: "E"} -> ["", "E", ""]
	// arr := [...]string{1: "E"}

	const (
		TH int = iota
		EN
		CH
	)
	capitals := [...]string{
		TH: "Bangkok",
		EN: "London",
		CH: "Japan",
	}
	fmt.Println(capitals[1])
	fmt.Println(capitals[EN])
}
