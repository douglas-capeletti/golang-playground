package main

import (
	"fmt"
	"github.com/going/concorrencia/generator"
	"time"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := generator.Titulo(url1)
	c2 := generator.Titulo(url2)
	c3 := generator.Titulo(url3)

	// estrutura de controle específica para concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
		// default:
		// 	return "Sem resposta ainda!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.google.com",
	)
	fmt.Println(campeao)
}
