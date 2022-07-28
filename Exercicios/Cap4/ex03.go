//- Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
// - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
// - Do quinto ao último item do slice (incluindo o último item!)
// - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
// - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
// - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item

package main

import "fmt"

func main () {
		slice := []int {10,20,30,40,50,60,70,80,90,100}
		fmt.Println(slice)

		l := slice[:3]
		fmt.Println("Do primeiro ao terceiro item do slice:\n", l)

		s := slice [4:]
		fmt.Println("Do quinto ao último item do slice:\n",s)

		a:= slice [1:7]
		fmt.Println("Do segundo ao sétimo item do slice:\n",a)

		b:= slice [2:9]
		fmt.Println("Do terceiro ao penúltimo item do slice:\n", b)

		//Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
		fmt.Println("\n\n",slice [2:len(slice)-1])
		

		/* Outra forma de fazer= apenas mostrando com fmt sem criar variáveis 
			fmt.Println(slice[:3])
			fmt.Println(slice[4:])
			fmt.Println(slice[1:7])
			fmt.Println(slice[2:9])
		*/

		}