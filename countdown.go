// Listing 3.8
package main

import (
	"fmt"
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
 	
}