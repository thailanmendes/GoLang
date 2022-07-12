// Atribua um valor int a uma variável
// Demonstre este valor em decimal, binário e hexadecimal

package main

import "fmt"

func main () {
	a := 200
	b := a << 1
	fmt.Printf(" %d\t %b\t %#x\n",a,a,a)
	fmt.Printf(" %d\t %b\t %#x\n",b,b,b)
}