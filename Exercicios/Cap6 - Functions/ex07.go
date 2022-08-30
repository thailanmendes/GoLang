//- Atribua uma função a uma variável.
//- Chame essa função.

package main

import "fmt"

func main () {
	x := 42	
	y:= func(x int) int {
		return x * 10
	}
	fmt.Println(x, "vezes 10 é:", y(x))
}