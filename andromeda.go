// Listing 8.2 (days to Andromeda Galaxy)
// big.Int is for big integers, when 18 quintillion isn’t enough.
// big.Float is for arbitrary-precision floating-point numbers.
// big.Rat is for fractions like ⅓.
package main

import (
	"fmt"
	"math/big"
)

func main()  {
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)

	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)
	fmt.Println("Andromeda Galaxy is", distance, "km away.")

	seconds :=new(big.Int)
	seconds.Div(distance, lightSpeed)
	
	days :=new(big.Int)
	days.Div(seconds, secondsPerDay)
	fmt.Println("That is", days, "days of travel at light speed.")
 	
}