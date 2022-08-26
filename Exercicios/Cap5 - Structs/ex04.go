//- Crie e use um struct anônimo.
//- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.

package main

import "fmt"

func main () {
	x := struct {
		CopaMundo string
		Cidade string
		Esportes []string           //tipo slice
		NotaEsporte map[string]float64  //tipo map
	}{
		CopaMundo: "Brasil",
		Cidade: "Brasilia",
		Esportes: []string{"Futebol","Basquete"},
		NotaEsporte: map[string]float64{
			"futebol": 9.99,
			"basquete": 8.5,
		},
}

	for _, valor := range x.Esportes{
		fmt.Println(valor)
	}

	fmt.Println(x)


}