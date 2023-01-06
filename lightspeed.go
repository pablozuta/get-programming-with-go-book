// 2.3 Page 17
// How long does it take to get to Mars?
// Light travels at a constant speed in the vacuum of space, which makes the math easy. On the other hand, the distance between Earth and Mars varies significantly, depending on where the planets are in their orbits around the Sun.
package main

import "fmt"

func main() {
	const lightSpeed = 299792 // kilometros por segundo
	var distance = 56000000 // kilometros
	fmt.Println(distance/lightSpeed, " seconds")

	distance = 401000000
	fmt.Println(distance/lightSpeed, " seconds")

	// SpaceX Interplanetary Transport System
	var spaceVelocidad = 100800 // Kilometros por hora
	var spaceDistancia = 96300000 // distancia a marte
	const horasDia = 24
	
	fmt.Print("El viaje de SpaceX a Marte tomara: ")
	fmt.Println(spaceDistancia/spaceVelocidad/horasDia, " Dias")

}