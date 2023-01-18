// Listing 3.1 , Page 24
package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println("You find yourself in a dim cavern.")
	var command = "walk outside"
	
	// Several functions in the standard library return a Boolean value. 
	// For example, the following listing uses the Contains function from 
	// the strings package to ask if the command variable contains 
	// the text “outside”. It does contain that text, so the result is true.
	var exit = strings.Contains(command, "outside")
	fmt.Println("You leave the cave:", exit)

	var wearShades = true
	fmt.Println(wearShades)

	var sign = "please read this sign"
	var signReading = strings.Contains(sign, "read")
	fmt.Println("you read the sign", signReading)
 
	
}