package main

import "fmt"

func main() {

	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	// polimorfismo
	coisa = produto{"Calça Jeans", 79.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	// não eh necessario estar no nivel da interface para acessar um metodo que necessita dela
	p2 := produto{"Calça Jeans", 179.90}

	// compilador reconhece que em produto é implementada a interface imprimivel
	imprimir(p2)

	//no nivel da interface eh necessario passar um ponteiro de carro
	//no metodo ligarTurbo() o ponteiro é necessário
	var carro1 esportivo = &carro{"F40", false, 0}
	carro1.ligarTurbo()

	//estando no nivel do carro o comportamente é padrao
	//por lidar diretamente com o carro a passagem de ponteiro é implicita
	carro2 := carro{"F40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro2, carro1)
}

// Cenario 1
type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

// qualquer struct que implemente a interface
func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

//Cenario 2
type esportivo interface {
	ligarTurbo()
}

type carro struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *carro) ligarTurbo() {
	f.turboLigado = true
}
