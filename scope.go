// In Go, scopes tend to begin and end along the lines of curly braces {}.
// In the following listing, the main function begins a scope
// and the for loop begins a nested scope.
package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	var count = 0

	for count < 10 { // a new scope begins
		var num = rand.Intn(10) + 1
		fmt.Println(num)
		count++
	} // this scope ends
 	
}