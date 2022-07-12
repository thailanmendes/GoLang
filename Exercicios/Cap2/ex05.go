// Crie uma variável de tipo string utilizando uma raw string literal ` `
// Demonstre-a

package main

import "fmt" 


func main () {
	name := ` isso é 
	
	raw   string                 literal
	
	(essas aspas   
		invertidas   )`
	fmt.Println(name)
}