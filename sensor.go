// Assigning functions to variables
// The weather station sensors provide an air temperature reading from 150–300º K.
// You have functions to convert Kelvin to other temperature units once you have the data 
// but unless you have a sensor attached to your computer (or Raspberry Pi)
// retrieving the data is a bit problematic.
// For now you can use a fake sensor that returns a pseudorandom number
// but then you need a way to use realSensor or fakeSensor interchangeably.
// The following listing does just that.
// By designing the program this way, different real sensors could also be plugged in,
// for example, to monitor both ground and air temperature.
package main

import (
	"fmt"
	"math/rand"
)
type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

func main()  {
 	sensor := fakeSensor // assign a function to a variable
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())
}