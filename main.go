package main

import "fmt"

func main() {
	// s1 := []int{1, 2, 3, 4, 5}
	// fmt.Println(s1, len(s1), cap(s1))

	// MAKE
	s2 := make([]int, 0, 5) //length = 1, capacity = 5
	s2 = append(s2, 1, 2)

	s3 := append(s2, 6)
	s3[0] = 10 // This will change both s2[0] and s3[0] to 10 because it shared the same array

	// BUT IF WE DO THIS
	// if s2 := make([]int, 0, 5)
	// s2 = append(s2, 1, 2, 3, 4, 5)
	// s3 := append(s2, 6)
	// this will make s2 = [1 2 3 4 5] and s3 = [1 2 3 4 5 6] which are different array because the s2 is already full (due to capacity of 5)

	// SO, IF WE DO
	// s3[0] = 10 -> This will change only s3[0] to 10

	fmt.Println(s2, s3)
}
