package main

import (
	"fmt"
)

func main() {
	arr := [5]string{"A", "B", "C", "D", "E"} // fixed size = 5

	// result := arr[0:3] // 3 is excluded (0 to 2) -> result = [A B C]
	// result is SLICE ([]string) not ARRAY ([x]string) -> x can be num or ...
	// array not fixed size, so it can expand
	// ** SLICE is more flexible than ARRAY

	// SLICE
	// s := arr[1:5]
	// STORE: address pointer, length, cap
	// pointer point to original array start position memory assress (arr[1]) -> 0x1111
	// len(s) -> length -> returns 4
	// cap(s) -> capacity -> returns 4

	// arr[0:5] == arr[:] -> get all from array to slice
	// arr[0:3] == arr[:3] -> ignore the first value means 0
	// arr[2:5] == arr[2:]  -> ignore the second value means 5

	// result[0] = "Z"
	// fmt.Println(result, arr) // [Z B C] [Z B C D E] -> the first item will change on both slice because they shared the same data

	result := arr[2:4] // [C D]
	result[0] = "M"    // [M D]

	fmt.Println(result, arr) // new arr will = [A B M D E]

	// CREATE SLICE WITHOUT CREATE ARRAY
	s := []string{"A", "B", "C", "D", "E"}
	// this will create array *in the background*
	// STORE -> address pointer, len = 5, cap =5

	// ADD MORE VALUE TO THE SLICE
	s2 := append(s, "F") // (target slice, new value)
	// this value will append to the related ARRAY
	// ARRAY capacity is fixed, so it will *create a new array* with a copied value + new value
	// ** GO is clever -> it create *more 5 empty slot* in the new array (to prevent too much re-creation)
	// s2 address pointer will point to the new array (len = 5,cap = 5)
	// s address point will point to the old array (len = 6, cap = 10)

	// So, THE DIFFERENCE BETWEEN LEN AND CAP is
	// len counts only slot filled with value -> the length of current member
	// cap counts all the slot (the empty slot is included) -> the available space for member
}
