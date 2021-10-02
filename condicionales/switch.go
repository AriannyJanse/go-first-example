package main

import "fmt"

func main() {
	fmt.Println("Hello, what is your language?")
	var language string
	_, err := fmt.Scanf("%s", &language)
	if err != nil {
		fmt.Println("Sorry, something went wrong")
		panic(err)
	}

	var name string

	switch language {
	case "es":
		fmt.Println("Ingresa tu nombre")
		_, err := fmt.Scanf("%s", &name)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Bienvenido %s, al maravilloso mundo de Golang", name)
	case "en":
		fmt.Println("Enter your name")
		_, err := fmt.Scanf("%s", &name)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Welcome %s, to the marvelous world of Golang", name)
	case "fr":
		fmt.Println("Entrez votre nom")
		_, err := fmt.Scanf("%s", &name)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Bienvenue %s, dans le monde merveilleux du Golang", name)
	case "de":
		fmt.Println("Gib deinen Namen ein")
		_, err := fmt.Scanf("%s", &name)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Willkommen %s, in der wunderbaren Welt von Golang", name)
	default:
		fmt.Println("Sorry your language is not supported yet")

	}
}
