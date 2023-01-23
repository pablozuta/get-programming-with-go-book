// 2.1 - 2.2 Page 14
// Mars takes 687 Earth days to travel around the sun
// the weight on Mars is 37.83% of the weight on Earth.

package main

import "fmt"

// main is the function when it all begins
func main() {
	fmt.Print("My weight on the surface of Mars is: ")
	// esto de aca multiplica el peso de la persona por el porcentaje en marte
	fmt.Print(72 * 0.3783)
	fmt.Print(" kilos, and I would be ")
	fmt.Print(40 * 365 / 687)
	fmt.Println(" Years old.")
	

	//podemos hacer esto mismo en una sola linea usando Println
	fmt.Println("Mi peso en la superficie de Marte seria: ", 72 * 0.3783, " Kilos , y mi edad serian los juveniles", 40 * 365 / 687, "años.")

	// age with input from person
	var inputAge int16
	fmt.Println("Please enter your age:")
	fmt.Scanln(&inputAge)
    var marsAge = inputAge * 365 / 687
	fmt.Println("Your age in martian years is: " , marsAge)
	
	// weight with input
	var inputWeight float32
	fmt.Println("Please enter your weight: ")
	fmt.Scanln(&inputWeight)
	var marsWeight = inputWeight * 0.3783
	fmt.Println("Your weight in martian surface would be : " , marsWeight)

}