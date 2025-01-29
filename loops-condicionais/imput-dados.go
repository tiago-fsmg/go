package main

import "fmt"

func main() {

	//Variaveis
	var nome string
	var idade int
	var altura float64
	var ativo bool

	fmt.Printf("Digite o nome:")
	fmt.Scanln(&nome)

	fmt.Printf("Gigite a idade(ex 30):")
	fmt.Scanln(&idade)

	fmt.Printf("Digite a altura(ex 1.75):")
	fmt.Scanln(&altura)

	fmt.Printf("Esta Ativo no sistema? (true/false):")
	fmt.Scanln(&ativo)

	fmt.Printf("\nNome: %s\nIdade: %d anos \nAltura: %.2f metros \nAtivo?: %t\n", nome, idade, altura, ativo)

}
