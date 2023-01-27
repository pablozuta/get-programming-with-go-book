// Listing 7.2 (Integers wrap around)
// Integers are free of the rounding errors that make floating-point inaccurate
// but all integer types have a different problem: a limited range. 
// When that range is exceeded, integer types in Go wrap around.
package main

import (
	"fmt"
)

func main()  {
	var red uint8 = 255
	red++
	fmt.Println(red) // Prints 0

	var number int8 = 127
	number++
	fmt.Println(number) // Prints -128
 	
}