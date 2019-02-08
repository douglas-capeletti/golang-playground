package main

import "fmt"

func main() {
	fmt.Println(soma(2, 3))
	sub(2, 3)

	// chama a funcao passando outra como parametro
	fmt.Println(exec(multiplicacao, 3, 4))
}

// funcoes aninhadas
func sub(val1, val2 int) {
	sub := func(val1, val2 int) int {
		return val1 - val2
	}
	fmt.Println(sub(val1, val2))
}

// funcao armazenada em uma variaval
var soma = func(a, b int) int {
	return a + b
}

// funcao como parametro
// a assinatura completa da funcao necessaria deve estar declarada nos parametros
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

// exemplo
func multiplicacao(a, b int) int {
	return a * b
}
