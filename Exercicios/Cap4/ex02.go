//- Usando uma literal composta:
// - Crie uma slice de tipo int
//- Atribua 10 valores a ela
//- Utilize range para demonstrar todos estes valores.
//- E utilize format printing para demonstrar seu tipo.

package main

import "fmt"

func main () {
		slice := []int {10,20,30,40,50,60,70,80,90,100}
		fmt.Println(slice)

		for indice, valor := range slice{
			fmt.Println("indice:",indice,"valor:",valor)
		}
			fmt.Printf("Tipo do slice = %T",slice)
}