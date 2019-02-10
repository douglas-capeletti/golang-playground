package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// pode adicionar mais métodos
}

type carro struct{}

func (b carro) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b carro) fazerBaliza() {
	fmt.Println("Baliza...")
}

func main() {
	var b esportivoLuxuoso = carro{}
	b.ligarTurbo()
	b.fazerBaliza()
}
