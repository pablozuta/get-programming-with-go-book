// Every element of an array has the same type.
// In this case, planets is an array of strings.
// Individual elements of an array can be accessed by using square brackets []
// with an index that begins at 0.
package main

import (
	"fmt"
)

func main()  {
 	var planets[8] string
	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"
	planets[3] = "Saturn"
	planets[4] = "Jupiter"
	planets[5] = "Uranus"

	fmt.Println(planets[0])
	fmt.Println(planets)

	
	// The length of an array can be determined with the built-in len function.
	// The other elements contain the zero value for their type, an empty string:
	fmt.Println("Lenght of planets array= ", len(planets))
}