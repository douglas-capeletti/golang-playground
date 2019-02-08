package main

import "fmt"

func init() {
	fmt.Println("Inicializando...")
}

func main() {
	//todos os init(s) são executados antes da funcao main
	fmt.Println("Main...")
}
