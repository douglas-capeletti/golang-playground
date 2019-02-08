package main

import "fmt"

func main() {
	x := 20
	fmt.Println(x)

	//toda funcao em go respeita o conceito de closure
	imprimeX := closure()
	imprimeX()
}

//closure mantem o contexto no qual a variavel eh declarada
func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}
