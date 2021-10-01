package main

import "fmt"

func main() {
	fmt.Println("Ingresa un número")

	var numero int

	_, err := fmt.Scanf("%d", &numero)
	if err != nil {
		panic(err)
	}

	if numero > 0 {
		fmt.Println("El número es mayor que 0")
	} else if numero == 0 {
		fmt.Println("El número es igual a 0")
	} else {
		fmt.Println("El número es menor que 0")
	}
}
