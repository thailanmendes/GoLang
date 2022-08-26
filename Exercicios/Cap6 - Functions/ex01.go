//- Crie uma função que retorne um int
//- Crie outra função que retorne um int e uma string
//- Chame as duas funções
//- Demonstre seus resultados

package main

import "fmt"

func main () {
	valor := retornaint(100)
	fmt.Println(valor)
	fmt.Println(retornaintstring(10, "dez"))

}

func retornaint (x int) int {
		return x

}

func retornaintstring (a int, b string) (int,string){
		return a,b

}