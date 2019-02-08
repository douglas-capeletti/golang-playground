package main

import "fmt"

func main() {
	//Sem retorno
	f1()
	f2("Param1", "Param2")

	//retorno unico
	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Println(r3)
	fmt.Println(r4)

	//multiplo retorno
	r51, r52 := f5()

	//Ao capturar o retorno de uma funcao deve-se capturar todos os retorno
	// r := f5() // codigo com erro

	//pode-se capturar o elemento que interessa e ignorar os restantes com underline
	// r1, _ := f5() // codigo correto
	fmt.Println("F5:", r51, r52)
}

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}
