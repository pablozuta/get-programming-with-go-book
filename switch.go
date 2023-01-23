// Or you can use the switch statement with conditions
// for each case, much like using if…else.
// One unique feature of switch is the fallthrough keyword
// which is used to execute the body of the next case
package main

import (
	"fmt"
)

func main()  {
	var room = "lake"

	switch {
		// expressions are in each case
	case room == "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case room == "lake":
		fmt.Println("The ice seems solid enough")
		// falls to the next case
		fallthrough
	case room == "underwater":
		fmt.Println("The water is freezing cold")
	}
 	
}