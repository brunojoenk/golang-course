package main

import "fmt"

func main() {
	i := 1 //ocupa 64 bytes na memoria

	//ver programação orientada a local
	var p *int = nil //endereço da variavel

	p = &i //pegando o endereço da variavel
	*p++   //desreferenciando o ponteiro (pegando o valor)
	i++

	//go não tem aritmética de ponteiros,
	//mas pode compartilhar com várias referencias
	//para ser utilizado pelas variáveis

	fmt.Println(p, *p, i, &i)
}
