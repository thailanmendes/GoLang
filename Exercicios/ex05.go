package main

import "fmt"

type thailan int
var x thailan
var y int

func main () {
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n",y)
}