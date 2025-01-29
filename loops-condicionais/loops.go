package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//Loop FOR Simples
	fmt.Println("Impriminto Tabuada")

	for i := 1; i <= 10; i++ {
		fmt.Println("")
		for j := 1; j <= 10; j++ {

			resultado := j * i
			fmt.Println(i, "x", j, " = ", resultado)
		}
	}

	//Loop FOR como whille
	rand.Seed(time.Now().UnixNano()) //Inicia a aleatoriedade
	num := 0

	fmt.Println("")
	for num < 90 {
		num = rand.Intn(100) //Gera um Numero entre 0 e 99
		fmt.Println("Numero Gerado: ", num)

	}

	fmt.Println("Numero Maior que 90 encontrado!")

	//Loop FOR Range
	numeros := []int{10, 20, 30, 40, 50, 60}

	for indice, valor := range numeros {
		fmt.Printf("Indice: %d, Valor: %d\n", indice, valor)
	}

	//Loop FOR Range Mapas
	pessoas := map[string]int{"Alice": 25, "Bob": 30, "Carlos": 35}

	for nome, idade := range pessoas {
		fmt.Printf("%s tem %d anos\n", nome, idade)
	}

}
