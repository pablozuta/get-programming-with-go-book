// You can assign an anonymous function to a variable and then
// use that variable like any other function,
package main

import (
	"fmt"
)

// assigns an anonymous function to a variable
var f = func() {
	fmt.Println("Dress up for the masquerade.")
}

func main()  {
	// print "Dress up for the masquerade"
 	f()
}