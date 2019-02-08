package main

import "fmt"

func main() {
	resultado, _ := fatorialVerificacaoErro(5)
	fmt.Println(resultado)

	_, err := fatorialVerificacaoErro(-4)
	if err != nil {
		fmt.Println(err)
	}

	resultado2 := fatorialSemSinal(5)
	fmt.Println(resultado2)
}

func fatorialVerificacaoErro(num int) (int, error) {
	switch {
	case num < 0:
		return -1, fmt.Errorf("número inválido: %d", num)
	case num == 0:
		return 1, nil
	default:
		numAnterior, _ := fatorialVerificacaoErro(num - 1)
		return num * numAnterior, nil
	}
}

func fatorialSemSinal(num uint) uint {
	switch {
	case num == 0:
		return 1
	default:
		return num * fatorialSemSinal(num-1)
	}
}
