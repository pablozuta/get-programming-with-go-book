// Go has two floating-point types. The default floating-point type is float64
// a 64-bit floating-point type that uses eight bytes of memory.

// The float32 type uses half the memory of float64 but offers less precision.

package main

import (
	"fmt"
	"math"
)

func main()  {
	var pi64 = math.Pi
	var pi32 float32 = math.Pi

	fmt.Println(pi64)
	fmt.Println(pi32)
 	
}