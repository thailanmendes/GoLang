//- Crie um tipo struct "pessoa" que contenha os campos:
//    - nome
//    - sobrenome
//    - idade
//- Crie um método para "pessoa" que demonstre o nome completo e a idade;
//- Crie um valor de tipo "pessoa";
//- Utilize o método criado para demonstrar esse valor.

package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade int
}

func (p pessoa) Nomeinteiro(){
	fmt.Println("o nome completo dessa pessoa é", p.nome, p.sobrenome,"\n e essa pessoa tem ", p.idade, "anos de idade")
}

func main() {
	/*thaila := pessoa {
		nome : "Thailan",
		sobrenome : "Mendes",
		idade :     29,
	}*/
	thailan := pessoa{"Thailan","Mendes",29}  //maneira simplificada 
	thailan.Nomeinteiro()
}