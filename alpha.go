// Listing 8.1 (Days to Alpha Centauri)
// If you haven’t realized it yet, 64-bit integers are mind-bogglingly big
// much bigger than their 32-bit counterparts.

// For some perspective, the nearest star, Alpha Centauri
// is 41.3 trillion kilometers away.
// A trillion: that’s one followed by 12 zeros, or 1012. 
// Rather than painstakingly typing every zero
// you can write such numbers in Go with an exponent, like so:
// var distance int64 = 41.3e12
package main

import (
	"fmt"
)

func main()  {
	const lightSpeed = 299792 // km/s
	const secondsPerDay = 86400

	var distance int64 = 41.3e12
	fmt.Println("Alpha Centauri is", distance, "km away.")

	days := distance / lightSpeed / secondsPerDay
	fmt.Println("That is", days, "days travel al light speed.")
 	
}