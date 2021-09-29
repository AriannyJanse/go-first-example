package main

import "fmt"

// La función main es el Entry Point de nuestro programa
func main() {
	// Declaración de una variable
	var name string

	fmt.Printf("¿Cuál es tu nombre? \n")

	/* Capturar por consola la respuesta del usuario y
	almacenarla en la variable name
	*/
	fmt.Scanf("%s", &name)

	// Imprimir mensaje utilizando una variable
	fmt.Printf("Hola %s, bienvenido al maravilloso mundo de Golang", name)

	//fmt.Println("Hola Mundo desde Golang")
}
