package main

import "fmt"

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}

func obterValorAprovado(numero int) int {

	//atrasa a execucao para o momento exato antes do retorno
	defer fmt.Println("fim!")

	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero
}
