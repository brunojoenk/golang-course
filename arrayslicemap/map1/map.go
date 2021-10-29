package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[1] = "Maria"
	aprovados[2] = "João"
	aprovados[3] = "Ana"
	//fmt.Println(aprovados)

	/*for chave, nome := range aprovados {
		fmt.Printf("%s (Chave: %d)\n", nome, chave)
	}*/

	//fmt.Println(aprovados[2])
	delete(aprovados, 2)
	fmt.Println(aprovados[2])

}
