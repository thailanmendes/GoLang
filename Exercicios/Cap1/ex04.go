package main

import "fmt"

type thailan int
var x thailan

func main () {
	fmt.Printf("%v\n%T",x,x)
	x = 42
	fmt.Println(x)
}