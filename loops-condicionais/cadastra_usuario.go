package main

import "fmt"

// Declara Parametros da estrutura
type Usuario struct {
	Nome   string
	Idade  int
	Altura float64
	Ativo  bool
}

func main() {
	//Declara Variaveis
	var usuarios []Usuario

	for {
		var nome string
		var idade int
		var altura float64
		var ativo bool

		fmt.Print("Digite o nome (ou 'sair' para finalizar): ")
		fmt.Scanln(&nome)

		if nome == "sair" {
			break
		}

		fmt.Print("Digite a idade(ex: 30): ")
		fmt.Scanln(&idade)

		fmt.Print("Digite a altura(ex: 1.75): ")
		fmt.Scanln(&altura)

		fmt.Print("Esta Ativo? (true/false): ")
		fmt.Scanln(&ativo)

		usuario := Usuario{Nome: nome, Idade: idade, Altura: altura, Ativo: ativo}
		usuarios = append(usuarios, usuario)

		fmt.Println("Usuario Cadastrado!\n")
	}

	fmt.Println("\nLista de Usuarios Cadastrados:")
	for _, u := range usuarios {
		fmt.Printf("Nome:	%s	| Idade:	%d	| Altura:	%.2f	| Ativo:	%t\n", u.Nome, u.Idade, u.Altura, u.Ativo)
	}

}
