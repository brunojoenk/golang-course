package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	//fmt.Println("Executou!") Após aqui o buffer encheu, só vai enviar o quarto elemento quando os outros forem consumidos
	ch <- 4
	fmt.Println("Executou!") // Aqui pode ser que apareça no log porque pode ser que passe o 1 segundo e libere o espaço o buffer, ja que o valor é lido
	ch <- 5                  // A partir daqui seria impossivel aparecer, porque só tiramos um item do buffer
	ch <- 6
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
