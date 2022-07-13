package main

import "fmt"

func main () {
	for i := 65; i <= 90; i++{
		fmt.Println(i)
		for u := 1; u <=3; u++ {
			fmt.Printf("\t%#U \n", i)
		}
	}
	
}