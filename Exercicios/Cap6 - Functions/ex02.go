//- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
//- Passe um valor do tipo slice of int como argumento para a função;
//- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
//- Passe um valor do tipo slice of int como argumento para a função.

package main

import "fmt"

func main () {
		//fmt.Println(soma(10,10,10))


		slice :=[]int {10,20,30,40,50}
		fmt.Println(soma(slice...))

		slice2 := []int {10,20,30,40,50,60,70,80,90,100}
		fmt.Println(somaslice(slice2))
}

func soma (x ...int) int {   
		soma := 0
		for _,v := range x {
			soma += v
		}
		return soma
} 

func somaslice (a []int) int{
		total := 0
		for _,v := range a {
			total += v
		}
		return total

}

