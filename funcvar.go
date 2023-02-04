// The variable you declare can be in the scope of the package or within a function,
package main

import (
	"fmt"
)

func main()  {
	// Assigns an anonymous function to a variable
	f := func(message string) {
		fmt.Println(message)
	}
	// Prints Go to the party.
 	f("Go to the party.")
}