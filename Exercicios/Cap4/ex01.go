//- Usando uma literal composta:
//- Crie um array que suporte 5 valores to tipo int
//- Atribua valores aos seus índices
//- Utilize range e demonstre os valores do array.
//- Utilizando format printing, demonstre o tipo do array.

package main

import "fmt"



func main () {
		array := [5]int{10,20,30,40,50}
		fmt.Println(array)

		for indice, valor := range array{
			fmt.Println("No indice",indice, "temos o valor =",valor)
		}

		fmt.Printf("O tipo do array é : %T", array)

}