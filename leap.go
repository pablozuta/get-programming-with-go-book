// The code in listing 3.4 determines whether 2100 will be a leap year.
// Any year that is evenly divisible by 4 but not evenly divisible by 100
// Or any year that is evenly divisible by 400
package main

import (
	"fmt"
)

func main()  {
	fmt.Println("The year is 2100, should you leap?")

	var year = 2100
	// || significa "or" && significa "and"
	var leap = year % 400 == 0 || (year % 4 == 0 && year % 100 != 0)

	if leap {
		fmt.Println("Look before you leap!")
	} else {
		fmt.Println("Keep your feet on the ground.")
	}
  	
}