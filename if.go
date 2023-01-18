// Listing 3.3
// A computer can use Boolean values or comparisons to choose
// between different paths with an if statement, as shown in the following listing.
// P-26
package main

import (
	"fmt"
)

func main()  {
	var command = "go east"

	if command == "go east" {
		fmt.Println("You head further up the mountain.")
	} else if command == "go inside" {
		fmt.Println("You enter the cave where you live out the rest of your life.")
	} else {
		fmt.Println("Didn't quite get that.")
	}

	// Quick check 3.3
	// Write a program that uses if and else if to display a description
	// for each of three rooms: cave, entrance, and mountain.
	var room = "cave"

	if room == "cave" {
		fmt.Println("This is the cave , the ultimate dungeon site to rest")
	} else if room == "entrance" {
		fmt.Println("You are now in the entrance to another magical place ")
	} else {
		fmt.Println("You are not in the cave or the entrance , you are up the mountain")
	}
 	
}