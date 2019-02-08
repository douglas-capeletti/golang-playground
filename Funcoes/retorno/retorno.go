package main

import "fmt"

func main() {
	x, y := trocar(5, 6)
	fmt.Println(x, y)
	x, y = trocar2(3, 4)
	fmt.Println(x, y)
	x, y = trocar3(1, 2)
	fmt.Println(x, y)
	x, y = trocar4(1, 2)
	fmt.Println(x, y)
}

//Retorno nomeado
//ao nomear o retorno as variaveis ja sao declaradas
func trocar(p1, p2 int) (segundo, primeiro int) {
	//atribuicao nas variaveis do retorno nomeado
	segundo = p2
	primeiro = p1

	//ao nomear o retorno nao eh necessario explicitar as variaveis no retorno
	return // retorno limpo
}

func trocar2(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return segundo, primeiro //mesmo que nao necessario as variaveis podem ser postas no retorno
}

func trocar3(p1, p2 int) (segundo, primeiro int) { //retorno com erro
	segundo = p2
	primeiro = p1
	return primeiro, segundo //retorno pode ser com ordem diferente do retorno declarado
}

func trocar4(p1, p2 int) (segundo, primeiro int) {
	return p2, p1
	//pode ser usado o retorno com variaveis diferentes das nomeadas (por mais que nao faca muito sentido nesse caso)
}
