//- Crie e utilize uma função anônima.

package main

import "fmt"

func main () {
	x := 20
	func(x int){
		fmt.Println(x,"x 95 é :", x * 95)
	}(x)
}