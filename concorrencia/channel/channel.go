package main

import (
	"fmt"
	"math"
)

func main() {

	simples()

	// Channel (canal) - é a forma de comunicação entre goroutines
	// eh um sincronizador entre dados das goroutines
	comGoroutine()
}

func simples() {

	//criando o canal (com buffer 1)
	// channel é um tipo
	ch := make(chan int, 1)

	// enviando dados para o canal (escrita)
	ch <- 1

	// recebendo dados do canal (leitura)
	valor := <-ch
	fmt.Println(valor)
}

func comGoroutine() {

	c := make(chan uint)
	go potenciador(2, 3, c)

	primeiro, segundo, terceiro := <-c, <-c, <-c // recebendo dados do canal
	fmt.Println(primeiro, segundo, terceiro)

	// ao ter somente tres dados no canal se o mesmo não for fechado ocorrera um DEADLOCK!
	fmt.Println(<-c)
}

func potenciador(x, y uint, c chan uint) {

	for i := 1; i <= int(y); i++ {
		c <- uint(math.Pow(float64(x), float64(i)))
	}
	//se comentado DEADLOCK!
	close(c)
}
