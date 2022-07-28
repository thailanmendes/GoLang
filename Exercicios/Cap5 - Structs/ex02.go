//- Utilizando a solução anterior, coloque os valores do tipo "pessoa" //num map, utilizando os sobrenomes como key.
//- Demonstre os valores do map utilizando range.
//- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome      string
	sabores   []string

}

func main () {
//nesse caso tem q ser feito com make
meumapa := make(map[string]pessoa)


//outra forma de fazer ...

/*meumapa2:= map[string]pessoa{
		"flocos": pessoa{
			nome: "Thailan",
			sobrenome: "Mendes",
			sabores: []string{"chocolate","flocos","pave"},
		},
		"morango": pessoa{
			pessoa{"Ruth","Raele", []string{"morango","prestigio"},
		},
}
*/

meumapa["Mendes"] = pessoa{
		nome: "Thailan",
		sobrenome: "Mendes",
		sabores: []string{"chocolate","flocos","pave"},
	}
meumapa["Raele"] = pessoa{"Ruth","Raele", []string{"morango","prestigio"},}

		for _,valor := range meumapa{
			fmt.Println("Meu nome é",valor.nome, "e meus sorvetes favoritos são:")
			for _, valor := range valor.sabores {
				fmt.Println("\t-",valor)
			}
		}
}
