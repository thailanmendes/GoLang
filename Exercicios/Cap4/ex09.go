//- Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.

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

		names["Trajano_José"] = []string{
			"Cortar Cabelo", "Sair a noite", "Dar banho no bob",
		}

	for key,valor := range names{
		fmt.Println(key)
		for indice , hobbie := range valor{
			fmt.Println("\t", indice, " - ", hobbie)
		}
	}
}


