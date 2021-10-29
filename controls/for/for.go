package main

import "fmt"

func main() {
	i := 1
	for i <= 10 { //nao tem while em go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando...")

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par, ", i)
		} else {
			fmt.Println("Impar, ", i)
		}
	}

	//for {
	//laço infinito
	//}
}
