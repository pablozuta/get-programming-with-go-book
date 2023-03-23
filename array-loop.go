package main

import (
	"fmt"
)

func main()  {
	dwarfs :=[5]string {"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	// muestra index y nombre planeta en la misma linea
	for i := 0; i < len(dwarfs); i++ {
		dwarf := dwarfs[i]
		fmt.Println(i, dwarf)
	}
 	
}