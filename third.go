// Listing 6.3 (Formatted print for floating points)
// When using Print or Println with floating-point types
//
//	the default behavior is to display as many digits as possible.
//
// If that’s not what you want, you can use Printf with the %f formatting
// verb to specify the number of digits, as the following listing shows.
package main

import (
	"fmt"
)

func main()  {
	third := 1.0 / 3
	fmt.Println(third) // Prints 0.333333333333

	fmt.Printf("%f\n", third) // Prints 0.333333
	fmt.Printf("%.3f\n", third) // Prints 0.333
	fmt.Printf("%4.2f\n", third) // Prints 0.33

	// The %f verb formats the value of third with a width and with precision

    // Precision specifies how many digits should appear after the decimal point
	// two digits for %.2f, for example

	// Width specifies the minimum number of characters to display
	// including the decimal point and the digits before and after the decimal 
	// (for example, 0.33 has a width of 4).
 	
}