package main

import "fmt"

type jogador struct {
	name      string
	sobrenome string
	idade     int
}

type competidor struct {
	jogador    jogador
	time       string
	velocista  bool
	competicao string
}

func main() {
	jogador1 := jogador{
		name:      "Lionel",
		sobrenome: "Messi",
		idade:     30,
	}

	jogador2 := competidor{
		jogador: jogador{
		name:      "Cristiano",
		sobrenome: "Ronaldo",
		idade:     32,
		},
		time:       "Manchester united",
		velocista:  true,
		competicao: "Premier",
	}
	fmt.Println(jogador1,jogador2)
}