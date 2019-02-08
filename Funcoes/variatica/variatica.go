package main

import "fmt"

func main() {
	fmt.Printf("Média: %.2f \n", media(7.7, 8.1, 5.9, 9.9))

	//slice
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}

	//ao colocar ... apos o slice, os elementos sao "espalhados" como parametros variaticos para a funcao
	imprimirAprovados(aprovados...)
}

func media(numeros ...float64) float64 {

	total := 0.0

	//parametros variaticos sao tratados como array dentro da funcao
	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}

func imprimirAprovados(aprovados ...string) {

	fmt.Println("Lista de Aprovados")

	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}
