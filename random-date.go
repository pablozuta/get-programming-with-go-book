// Listing 4.7
// To remove the duplication and simplify the code
// the variables in listing 4.6 should be declared in the wider function scope
// making them available after the switch statement for later work.
// It’s time to refactor!
// Refactoring means modifying the code without modifying the code’s behavior.
// The code in the following listing still displays a random date.
package main

import (
	"fmt"
	"math/rand"
)

var era = "AD" 

var month = rand.Intn(12) + 1
func main()  {
	daysInMonth := 31
	
	switch month {
	case 2:
		daysInMonth = 28
	case 4, 6, 9, 11:
		daysInMonth = 30	
	}
	day := rand.Intn(daysInMonth) + 1
	for count:= 0; count < 11; count ++ {
		var year = rand.Intn(2023) + 1
		fmt.Println(era, year, month, day)
			
	}
	

} 
// Though a narrower scope often reduces the mental overhead
// listing 4.6 demonstrates that constraining variables too tightly 
// can result in less readable code. 
// Take it on a caseby-case basis
// refactoring until you can’t improve the readability any further.