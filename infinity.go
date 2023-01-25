// An infinite loop doesn’t specify a for condition
// but you can still break out of a loop at any time.
// The following listing orbits a 360 circle and stops randomly.
package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	degrees := 0

	for {
		fmt.Println(degrees)
		degrees++
		if degrees >= 360 {
			degrees = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}
 	
}