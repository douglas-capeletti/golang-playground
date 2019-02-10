package main

import "fmt"

// Padrao de declaracao de tipo
type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {

	var produto1 produto

	// declaracao verbosa com nome de campos
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}

	//Metodos são accesivel atraves da notacao ponto
	fmt.Println(produto1, produto1.precoComDesconto())

	// declaracao simplificada, porém não tão clara
	produto2 := produto{"Notebook", 1989.90, 0.10}

	fmt.Println(produto2.precoComDesconto())
}
