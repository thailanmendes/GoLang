//- Crie um map com key tipo string e value tipo []string.
//- Key deve conter nomes no formato sobrenome_nome
//- Value deve conter os hobbies favoritos da pessoa
//- Demonstre todos esses valores e seus índices.

package main

import "fmt"

func main() {
	names := map[string][]string{
		"Mendes_Thailan": []string{
			"Jogar bola", "Videogame", "Academia", "estudar",
		},
		"Raele_Ruth": []string{
			"Assistir serie", "Novela", "ir a praia", "comer muito", "comida Japonesa",
		},
		"Mendes_Thauan" : []string {
			"Jogar VideoGame", "Game", "Brincar com o Pitu", 
		},
		"Cachorro_Bob" : []string{
			"Brincar com a bola", "Comer racao", "Tomar sol", "Passear na rua",
		},
	}
	for key,valor := range names{
		fmt.Println(key)
		for indice , hobbie := range valor{
			fmt.Println("\t", indice, " - ", hobbie)
		}
	}

}  //no map nao existe ordem fixa, é aleatória