package main

import (
	"fmt"
	"math/rand"
)

func main()  {
const numero = 10
var numeroRandom = rand.Intn(100)

for numero != numeroRandom {
	if numero < numeroRandom {
		fmt.Println("Numero Random es Mayor")
	}else {
		fmt.Println("Numero Random es Menor")
	}
numeroRandom = rand.Intn(100)
}
 fmt.Println("los numeros son iguales")	
}