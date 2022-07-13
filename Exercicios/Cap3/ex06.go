// Crie um programa que demonstre o funcionamento da declaração if.

package main

import "fmt"

func main () {
	vel := 101
	if vel <= 60 {
		fmt.Println("Parabens, esta andando conforme a velocidade permitida")
	}else if vel >= 80 && vel <= 100 {
		fmt.Println("Velocidade ultrapassada a permitida. Multado com 5 pontos ")
		}else{
			fmt.Println("Habilitacao cacada, Velocidade ultrapassada + de 20 km")
		}
			
	}
