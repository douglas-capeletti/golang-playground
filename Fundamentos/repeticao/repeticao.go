package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 { // não tem while em Go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando... ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	// O foreach pertence ao array
	numeros := []int{2, 3, 5, 7}

	//for range ignorando o indice e iterando sobre um array(slice)
	for _, num := range numeros {
		fmt.Println(num)
	}

	for {
		// laço infinito
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}

}
