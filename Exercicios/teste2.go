package main

import ("fmt"
		"time")

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type profissao struct {
	pessoa 
	atuacao string
	salario int
}

/*type sorvete struct {
	nome   string
	sabor  string
	lugar  []string
	vaibem map[string]string
}*/

type gente interface {
	oibomdia()
}

func interfac(g gente) {
	g.oibomdia()
}

func (p profissao) oibomdia() {
	fmt.Println("oi bom dia meu nome é", p.nome, "e tenho",p.idade,"anos de idade")
}

func main() {



	switch time.Now().Weekday() {
	case time.Friday, time.Sunday:
		fmt.Println("é final de semana")
	default:
		fmt.Println("é dia de semana ainda")
	}




	/*sabor := sorvete{
		nome:  "sorvete",
		sabor: "chocolate",
		lugar: []string{"Na Belgica", "é o melhor", "chocolate"},
		vaibem: map[string]string{
			"amargo":      "depois do almoco",
			"ao leite":    "de tardizinha",
			"meio amargo": "qualquer horaio",
		},
	}
	fmt.Println(sabor)*/


	thailan := profissao{
		pessoa: pessoa{
			nome:      "Thailan",
			sobrenome: "mendes",
			idade:     29,
		},
		atuacao: "motorista",
		salario: 2500,
	}

	fmt.Println(thailan)
	interfac(thailan)
	thailan.oibomdia()


}
