//- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
//- Nome
//- Sobrenome
//- Sabores favoritos de sorvete
//- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome      string
	sabores   []string

}

func main () {
	pessoa1 := pessoa{
		nome: "Thailan",
		sobrenome: "Mendes",
		sabores: []string{"chocolate","flocos","pave"},
	}
	pessoa2 := pessoa{"Ruth","Raele", []string{"morango","prestigio"},}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, valor := range pessoa1.sabores{
		fmt.Println("\t",valor)
	}

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, valor := range pessoa2.sabores{
		fmt.Println("\t",valor)
	}
}