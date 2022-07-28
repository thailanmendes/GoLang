//- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
//- "Nome", "Sobrenome", "Hobby favorito"
//- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dado

package main

import "fmt"

func main () {
		ss:= [][]string{
			[]string {"Thailan", "Mendes", "Jogar Bola"},
			[]string{"Ruth","Raele","Assistir Serie"},
			[]string {"Thauan","Mendes","Game"},
		}

		/*for _,valor := range ss {
			fmt.Prinln(valor)
		}*/
		for _,valor := range ss{
			fmt.Println(valor[0])
			for _, item := range valor{
				fmt.Println("\t",item)
			}
		}
		
		
}