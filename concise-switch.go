// When comparing one value to several others, Go provides the switch statement
package main

import (
	"fmt"
)

func main()  {
	fmt.Println("There is a cavern entrance here and a path to the east.")
	var command = "go inside"

	switch command {
	case "go east":
		fmt.Println("You head further up the mountain.")
	// A comma separated list of possible values	
	case "enter the cave", "go inside":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "read sign":
		fmt.Println("The sign reads 'No Minors'.")
	default:
		fmt.Println("Didn't quite get that.")		
	}
 	
}