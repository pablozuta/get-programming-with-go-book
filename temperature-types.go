package main

import (
	"fmt"
)
type celsius float64
type kelvin float64

// kelvinToCelsius converts ºK to ºC
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func main()  {
	// the argument must be of type kelvin
	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	fmt.Print(k, "º K is ", c, "ºC")
 	
}