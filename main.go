package main

import "fmt"

func main() {
	// var i int -> ZERO VALUE = 0
	// var i bool -> ZERO VALUE = false

	// var i int = 20

	// var i int
	// i = 20

	// i := 20 // NO NEED TO DECLARE TYPE
	// i = false -> NOT ALLOW

	// TYPE CONVERSION
	i := 20.02
	j := int(i) // j = 20 (int)

	// TYPE														ZERO VALUE
	// bool														false
	// string													""

	// int int8,16,32,64 -> Can store +/-/0						0
	// -> int8,16,32,64 = bit -> how much value can store
	// -> default int -> automatic matching with cpu bit (eg. 64 bit)

	// uint unint8,16,32,64 unitptr -> Can store + only			0

	// byte -> alias(nickname) for unit8
	// rune -> alias(nickname) for int32

	// float32,64 -> float num									0
	// ** no default float

	// complex64,128 -> complex num								(0 + 0i)

	fmt.Println(j)
}
