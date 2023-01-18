// The not logical operator (!) flips a Boolean value from false to true or vice versa.
// The following listing displays a message if the player doesn’t have a torch
// or if the torch isn’t lit.
// Listing 3.5 , P-29
package main

import (
	"fmt"
)

func main()  {
	var haveTorch = true
	var litTorch = false

	if !haveTorch || !litTorch {
		fmt.Println("Nothing to see here.")
	}
 	
}