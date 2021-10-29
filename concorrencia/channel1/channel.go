package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1   // enviando dados para canal (escrita)
	x := <-ch // recebendo dados do canal (leitura)
	fmt.Println(x)

	ch <- 2
	fmt.Println(<-ch)
}
