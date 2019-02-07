package main

import (
	"fmt"
)

func main() {
	declaracao()
	atribuicao()
	zeros()
	logicos()
}

func declaracao() {

	// Declaracao padrao
	const A float64 = 1
	var a float64 = 1

	// Tipos inferidos pelo compilador
	const B = 1
	var b = 1

	// Declaracao multipla com inferencia de tipos
	const (
		C = 1
		D = 2
	)
	var (
		c = 3
		d = 4
	)

	// Declaracao multipla inline com atribuicao
	const E, F bool = true, false
	var e, f bool = true, false

	// declaracao multipla em linha com inferencia de tipo na atribuicao
	// ao utilizar o ":" pode se omitiar o "var" que ja sera reconhecido pelo compilador, syntax sugar
	g, h, i := 2, false, "epa!"

	fmt.Println("Constantes:", A, B, C, D, E, F)
	fmt.Println("Variaveis:", a, b, c, d, e, f, g, h, i)
}

func atribuicao() {

	var b byte = 3
	fmt.Println(b)

	i := 3 // inferência de tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 2 // i = i / 2
	i *= 2 // i = i * 2
	i %= 2 // i = i % 2

	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)
}

func zeros() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int

	fmt.Printf("%v %v %v %q %v \n", a, b, c, d, e)
}

func logicos() {

	trab1, trab2 := true, false

	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // ou exclusivo
	comprarSorvete := trab1 || trab2

	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t",
		comprarTv50, comprarTv32, comprarSorvete, !comprarSorvete)
}
