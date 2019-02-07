package main

import (
	"fmt"
	"reflect"
)

func main() {
	sliceIntro()
	sliceMake()
	sliceCopy()
}

func sliceIntro() {

	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println("01)", a1, s1)
	fmt.Println("02)", reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	//Slice não é um array! Slice define um pedaço de um array.
	s2 := a2[1:3]
	fmt.Println("03)", a2, s2)

	s3 := a2[:2] // novo slice, mas aponta para o mesmo array
	fmt.Println("04)", a2, s3)

	//pode imaginar um slice como: tamanho e um ponteiro para um elemento de um array
	//atraves do tamanho o compilador sabe até espaço de memoria percorrer a partir do ponteiro
	s4 := s2[:1]
	fmt.Println("05)", s2, s4)

	//slice consegue efetuar codificacoes no array original conforme esperado de um ponteiro
	s5 := make([]int, 10, 20)
	s6 := append(s5, 1, 2, 3)
	fmt.Println("06)", s5, s6)

	s5[0] = 7
	fmt.Println("07)", s5, s6)
}

func sliceMake() {

	//cria um slice de tamanho 10
	s := make([]int, 10)

	//atruibuicao a um indice especifico
	s[9] = 12
	fmt.Println("08)", s)

	//slice de tamanho 10 e array no qual o slice aponta eh de tamanho 20
	s = make([]int, 10, 20)
	fmt.Println("09)", s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println("10)", s, len(s), cap(s))

	//ao atigir o tamanho maximo de um slice ele eh duplicado caso necessite de mais espaço
	s = append(s, 1)
	fmt.Println("11)", s, len(s), cap(s))
}

func sliceCopy() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 4, 5, 6)
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println("12)", array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println("13)", slice2)
}
