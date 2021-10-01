package main

import "fmt"

func main() {
	// For infinito con una sola ejecución
	for {
		fmt.Println("Infinito")
		break
	}

	veces := 1
	for {
		fmt.Println("Infinito ", veces)
		if veces == 5 {
			break
		}

		// veces = veces + 1 es igual a veces += 1
		veces++
	}
}
