package main

import (
	"fmt"
)

func main()  {
 	fmt.Println("There is a sign near the entrance that reads 'No Minors'.")
	var age = 41
	var minor = age < 18
	fmt.Printf("At age %v, am I a minor? %v\n", age, minor)

	// Go uses this %v as a placeholder?
	// yes , using the Printf command tho
	var nombre = "Pedro Paramo"
	fmt.Printf("Juan Rulfo escribio %v", nombre)
}