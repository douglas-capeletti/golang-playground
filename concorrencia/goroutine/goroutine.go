package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	fmt.Println("Exemplo 1):")
	ex1()
	fmt.Println("\nExemplo 2):")
	ex2()
	fmt.Println("\nExemplo 3):")
	ex3()
}

func ex1() {
	// funcao sendo executada de maneira sincrona
	fale("Maria", "Pq vc não fala comigo?", 3)
	fale("João", "Só posso falar depois de vc!", 1)
}

func ex2() {
	// ao inserir "go" a funcao eh chamada a partir de uma goroutine
	// uma ou varias goroutines sao executadas em uma thread separada do programa principal
	// se não tiver nada que faça a goroutine "esperar" o programa é finalizado
	go fale("Maria", "Ei...", 500)
	go fale("João", "Opa...", 500)
}

func ex3() {
	// somente será executado a primeira linha enquanto a execucao da segunda estiver ativa
	go fale("Maria", "Entendi!!!", 10)
	fale("João", "Parabéns!", 5)
}
