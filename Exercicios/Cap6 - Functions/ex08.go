//- Crie uma função que retorna uma função.
//- Atribua a função retornada a uma variável.
//- Chame a função retornada.

package main

import "fmt"

func main () {
	x := retornafunc()
	y := x(3)
	fmt.Println(y)
	
	//outra maneira de declarar 
	fmt.Println(retornafunc()(5))
}

func retornafunc() func(int)int{
		return func(i int)int {
			return i * 15
		}
}

