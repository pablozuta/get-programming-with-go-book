// Listing 6.2 (Declaring a variable without a value)
// In Go, each type has a default value, called the zero value.
// The default applies when you declare a variable but don’t initialize it with a value,

package main

import (
	"fmt"
)

func main()  {
	var price float64
	fmt.Println(price)
 	
}