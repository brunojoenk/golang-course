package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} //compilador conta

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for i := range numeros { //acessando o indice
		fmt.Println(i)
	}

	for _, numero := range numeros { //acessando o valor
		fmt.Println(numero)
	}
}
