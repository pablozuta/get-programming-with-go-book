// Listing 3.8
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	var cuenta = 10

	for cuenta > 0 {
		fmt.Println(cuenta)
		time.Sleep(time.Second)
		cuenta--
	}
	fmt.Println("Liftoff!")

	// random countdown
	for cuenta > 0 {
		fmt.Println(cuenta)
		time.Sleep(time.Second)
		if rand.Intn(100) == 0 {
			break
		}
	}
 	
}