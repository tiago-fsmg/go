package main

import "fmt"

func main() {

	//Variaveis
	idade := 25
	dia := "sexta"
	nota := 4

	//Condicional if-else
	if idade < 18 {
		fmt.Println("Idade: Menor de Idade")
	} else if idade == 18 {
		fmt.Println("Idade: Acabou de atingir a maioridade")
	} else if idade > 18 {
		fmt.Println("Idade: Maior de Idade")
	}

	//Switch com Numeros
	switch {
	case idade < 18:
		fmt.Println("Faixa: Menor de idade")
	case idade >= 18 && idade < 60:
		fmt.Println("Faixa: Adulto")
	default:
		fmt.Println("Faixa: Idoso")
	}

	//Switch com String
	switch dia {
	case "segunda":
		fmt.Println("Dia: Inicio da semana!")
	case "sexta":
		fmt.Println("Dia: Sextou!")
	case "sabado", "domingo":
		fmt.Println("Dia: Fim de semana!")
	default:
		fmt.Println("Dia: Comum")
	}

	//Switch com Fallthrough (Se quiser forçar a execução do próximo case, use fallthrough)

	switch {
	case nota >= 9:
		fmt.Println("Nota: Exelente!")
	case nota >= 7:
		fmt.Println("Nota: Aprovado!")
		fallthrough
	case nota >= 6:
		fmt.Println("Mas Pode Melhorar")
	default:
		fmt.Println("Nota: Reprovado")

	}
}
