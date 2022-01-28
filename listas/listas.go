package main

import "fmt"

func main() {
	// Las listas se dividen en dos: las listas mutables y las listas inmutables
	// Mutable significa que puede cambiar, en go se llaman slices
	// Inmutable significa que no puede cambiar, en go se llaman arrays

	// La mutabilidad o inmutabilidad no hace referencia a los valores sino al tamaño de la lista
	// El tamaño de una lista se define por cuántos elementos tiene

	// LISTAS INMUTABLES
	// Declaración de array
	var nombres [6]string

	nombres[0] = "Arianny"
	nombres[1] = "Naho"
	nombres[2] = "Sebas"
	nombres[3] = "David"
	nombres[4] = "Pepito"
	nombres[5] = "Zoro"

	fmt.Println(nombres)

	//nombres[5] = "Church"

	fmt.Println(nombres)
	fmt.Println("El valor de la posición 0:", nombres[0])

	// Declaración corta
	frutas := [4]string{}
	fmt.Println("El tamaño de la lista frutas es:", len(frutas))

	fmt.Println(frutas)

	notas := [5]int{5, 3, 4, 2, 1}

	fmt.Println(notas)
	fmt.Println("El valor de la posición 3 es:", notas[3])

	// Conocer el tamaño(length) de la lista
	fmt.Println("El tamaño de la lista nombres es:", len(nombres))

	// LISTAS MUTABLES
	// Declaración de slice
	var apellidos []string

	apellidos = append(apellidos, "Parra", "Janse")
	fmt.Println(apellidos)
	fmt.Println("El tamaño de apellidos es:", len(apellidos))

	apellidos = append(apellidos, "Parra", "Villa", "Perez", "Roronoa")
	fmt.Println("El tamaño de apellidos es:", len(apellidos))

	fmt.Println("El alumno de la posición 5 es:", nombres[5], apellidos[5])

	// Declaración corta
	edades := []int{23, 18, 15, 25, 30, 4}
	fmt.Println(edades)

	fmt.Println(nombres[0], apellidos[0], "tiene", edades[0], "años")

	// Recorriendo listas con for

	for contador := 0; contador < 6; contador++ {
		fmt.Println(nombres[contador])
	}

	for i := 0; i < len(apellidos); i++ {
		fmt.Println(apellidos[i])
	}

	// Recorriendo listas con for range

	for index, valor := range nombres {
		fmt.Println(valor, "está en el índice", index)
	}

	for _, value := range apellidos {
		fmt.Println(value)
	}

	for i, v := range nombres {
		fmt.Println(nombres[i])
		if v == "David" {
			fmt.Println("Hola David!!!")
			break
		}
	}

	for _, v := range nombres {
		if v == "David" {
			fmt.Println("Encontramos a David!!!")
			break
		}
	}

	fmt.Println("Ingresa el nombre que quieres buscar")
	var busqueda string
	_, err := fmt.Scanf("%s", &busqueda)
	if err != nil {
		fmt.Println("Ups! Algo salió mal")
		panic(err)
	}

	for i, v := range nombres {
		if v == busqueda {
			fmt.Println("El nombre fue encontrado con éxito!!!")
			break
		} else if i == len(nombres)-1 {
			fmt.Println("El nombre no fue encontrado😭")
		}
	}

}
