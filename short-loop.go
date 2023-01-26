// By using short declaration, the count variable in the next listing
// is declared and initialized as part of the for loop
// and falls out of scope once the loop ends.
// If count were accessed outside of the loop
// the Go compiler would report an undefined: count error.
package main

import "fmt"

func main()  {
	for count := 10; count > 0; count-- {
		fmt.Println(count)
	} // count variable is no longer in scope here
 	
}