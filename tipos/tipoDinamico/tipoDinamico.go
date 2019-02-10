package main

import "fmt"

type curso struct {
	nome string
}

func main() {

	//Uma interface vazia pode ser qualquer coisa
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	coisa = true
	fmt.Println(coisa)

	//pode ser criado um tipo de interface ao inves de somente uma variavel
	type dinamico interface{}

	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = false
	fmt.Println(coisa2)

	//tipo pode conter tambem uma struct
	coisa2 = curso{"Golang: Explorando a Linguagem do Google"}
	fmt.Println(coisa2)
}
