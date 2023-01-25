// The Print and Println functions have a sibling that gives more control over output. 
// By using Printf, shown in the following listing
// you can insert values anywhere in the text.
package main

import (
	"fmt"
)

func main()  {
	fmt.Printf("My weight on the surface of mars is %v lbs,", 149.0*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 41*365/687)
	// \n is used to jump to the next line

	// In addition to substituting values anywhere in a sentence
	// Printf can help you align text.
	// Specify a width as part of the format verb, such as %4v 
	// to pad a value to a width of 4 characters.
	// A positive number pads with spaces to the left
	// and a negative number pads with spaces to the right:
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n","Virgin Galactic", 100)
 	
}