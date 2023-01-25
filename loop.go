// Short declaration provides an alternative syntax for the var keyword.
// short declaration can go places where var can’t.
package main

import (
	"fmt"
)

func main()  {
	// using the var keyword
	var count = 0
	for count = 10; count > 0; count-- {
		fmt.Println(count)
	}
	fmt.Println(count) // count remains in scope
	fmt.Println("Termino del primer for loop")
	// Without short declaration, the count variable must be 
	// declared outside of the loop
	// which means it remains in scope after the loop ends.

	// By using short declaration, the count variable in the next listing
	// is declared and initialized as part of the for loop 
	// and falls out of scope once the loop ends.
	for countShort :=10; countShort > 0; countShort-- {
		fmt.Println(countShort)
	} // countShort is no longer in scope
 	
}