package main

import "fmt"

func main () {

	anoqnasci := 1993
	anoatual := 2022
	for {
		if anoqnasci > anoatual {
		break
	}
	  fmt.Println(anoqnasci)
	  anoqnasci++
	}
}