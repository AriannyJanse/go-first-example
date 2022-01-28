package main

import "fmt"

func main() {
	// Declarión de un mapcon la palabra reservada var
	// El tipo que va dentro de los corchetes son las claves, y el tipo que va fuera de los corchetes son los valores
	var mapa_numeros map[string]int

	// Inicializar mi map mediante la función make
	mapa_numeros = make(map[string]int)

	// Introducir valores a mi map
	mapa_numeros["uno"] = 1
	mapa_numeros["dos"] = 2
	mapa_numeros["tres"] = 3
	mapa_numeros["cuatro"] = 4

	fmt.Println(mapa_numeros)
	fmt.Println("El valor de la clave dos es:", mapa_numeros["dos"])

	// Declaración corta e inicialiazción
	nombres_apellidos := make(map[string]string)

	nombres_apellidos["Arianny"] = "Janse"
	nombres_apellidos["Naho"] = "Janse"
	nombres_apellidos["Sebas"] = "Parra"

	fmt.Println(nombres_apellidos)

	// Declarar, inicializar y añadir valores
	mascotas_duenos := map[string]string{
		"Ari":   "Zoro",
		"Naho":  "Sweet",
		"Sebas": "Church",
	}

	fmt.Println(mascotas_duenos)
	fmt.Println("La mascota de Ari es:", mascotas_duenos["Ari"])

	for i, v := range mascotas_duenos {
		fmt.Println(i, "es el/la dueño/a de", v)
	}
}
