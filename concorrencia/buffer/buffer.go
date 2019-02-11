package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// canal sem buffer
	ch := make(chan int)

	go random(ch)

	// operação random
	fmt.Println(<-ch)

	//assincronamente com o restante da funcao random()
	fmt.Println("Foi lido")
	fmt.Println("Fim")

	// canal com buffer
	chBuffer := make(chan int, 30)

	go primos(cap(chBuffer), chBuffer)

	for primo := range chBuffer {
		fmt.Printf("%d ", primo)
	}

	fmt.Println("Fim!")

}

func random(c chan int) {
	time.Sleep(time.Second)

	// operação random
	// valor retornado a goroutine
	c <- rand.Int()
	fmt.Println("Só depois que foi lido")
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 180)
				break
			}
		}
	}
	close(c)
}

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
