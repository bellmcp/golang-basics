package main

import "fmt"

func main() {
	// day := "Sunday"

	// switch day { // Case by case -> cannot group 2 case like other languages
	// case "Sunday":
	// 	fmt.Println("Weekend") // Auto break -> no need to break
	// case "Saturday":
	// 	fmt.Println("Weekend")
	// default:
	// 	fmt.Println("Workday")
	// }

	// // GROUPED CONDITION
	// switch day {
	// case "Sunday", "Saturday":
	// 	fmt.Println("Weekend")
	// default:
	// 	fmt.Println("Workday")
	// }

	// DECLARE AND CONDITION

	switch day := "Sunday"; day { // use ;
	case "Sunday", "Saturday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Workday")
	}
}
