// Listing 3.1 , Page 24
package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println("You find yourself in a dim cavern.")
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("You leave the cave:", exit)
 
	
}