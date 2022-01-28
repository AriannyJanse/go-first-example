package main

import "fmt"

func main() {
	type Persona struct {
		// Atributos
		Nombre           string
		Apellido         string
		Edad             int
		TieneMascota     bool
		ComidasFavoritas []string
	}

	type CuentaNequi struct {
		// Atributos
		Numero int
		Dueno  Persona
	}

	// Declaración de una variable de tipo persona con var
	var persona1 Persona
	fmt.Println(persona1)

	persona1.Nombre = "Arianny"
	persona1.Apellido = "Janse"
	persona1.Edad = 23
	persona1.TieneMascota = true
	persona1.ComidasFavoritas = []string{"Pasta", "Helado", "Camarones"}
	fmt.Println(persona1)

	fmt.Println(persona1.Nombre, "tiene mascotas? R//", persona1.TieneMascota)

	// Declaración con el operador de asignación corta
	persona2 := Persona{
		Nombre:           "Sebas",
		Apellido:         "Parra",
		Edad:             15,
		TieneMascota:     false,
		ComidasFavoritas: []string{"Sushi", "Arroz Mixto"},
	}

	fmt.Println(persona2)
	fmt.Println(persona2.Nombre, "tiene", persona2.Edad, "años")

	nueva_cuenta_nequi := CuentaNequi{
		Numero: 35554466,
		Dueno:  persona1,
	}
	fmt.Println(nueva_cuenta_nequi)

	otra_cuenta_nequi := CuentaNequi{
		Numero: 22223333000,
		Dueno: Persona{
			Nombre:           "Naho",
			Apellido:         "Janse",
			Edad:             18,
			TieneMascota:     true,
			ComidasFavoritas: []string{"Ramen", "Pollito frito"},
		},
	}
	fmt.Println(otra_cuenta_nequi)

	lista_personas := []Persona{
		persona1,
		persona2,
		Persona{
			Nombre:           "David",
			Apellido:         "Villa",
			Edad:             25,
			TieneMascota:     true,
			ComidasFavoritas: []string{"Pasta", "Arepa"},
		},
		Persona{
			Nombre:           "Amanda",
			Apellido:         "Janse",
			Edad:             42,
			TieneMascota:     false,
			ComidasFavoritas: []string{"Rabioes", "Empanadas"},
		},
	}

	fmt.Println(lista_personas)
}
