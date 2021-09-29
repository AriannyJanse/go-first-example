package main

import "fmt"

func main() {
	// Declaramos una variable llamada name
	var name string

	// Le asignamos un valor a la variable llamada name
	name = "Arianny"

	// Declaramos una variable llamada age y le asignamos un valor
	var age int = 23

	// Otra manera de declarar variables (Declaramos e inicializamos)
	hasPet := true

	// Imprimimos un mensaje de bienvenida con Printf
	fmt.Printf("Hola %s, tienes %d años, bienvenida al maravilloso mundo de Golang \n", name, age)

	// Conocer el tipo de una variable
	fmt.Printf("Tipo de name: %T \n", name)
	fmt.Printf("Tipo de age: %T \n", age)
	fmt.Printf("Tipo de hasPet: %T \n", hasPet)

	/*
		EL VALOR CERO

		string -> ""
		int -> 0
		float -> 0
		bool -> false
	*/

	var cadena string
	var entero int
	var decimal float32
	var logico bool

	fmt.Printf("Valor 0 de string: %v \n", cadena)
	fmt.Printf("Valor 0 de integer: %v \n", entero)
	fmt.Printf("Valor 0 de float: %v \n", decimal)
	fmt.Printf("Valor 0 de boolean: %v \n", logico)

}
