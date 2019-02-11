package main

import (
	"fmt"
	"github.com/going/concorrencia/generator"
)

func main() {
	c := juntar(
		generator.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		generator.Titulo("https://www.amazon.com", "https://www.youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}

// multiplexar - misturar (mensagens) num canal
// centraliza duas chamadas assincronas em um unico canal de retorno
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}
