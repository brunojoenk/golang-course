package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//integer numbers - Números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(3200))

	//sem sinal (só positivos)... uint8 uint16 uint32 uint64
	//u de unsigned - não tem sinal
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	//com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	// representa um mapeamento  da tabela Unicode (int32)
	//var i2 rune = 'a'
	//fmt.Println("O rune é", reflect.TypeOf(i2))
	//fmt.Println(i2)

	//números reais (float32, float64)
	var x = float32(49.99)
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do litireal 49.99 é", reflect.TypeOf(49.99)) // float64

	//boolean
	bo := true
	fmt.Println("O tipode bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//string
	s1 := "Olá meu nome é Bruno"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	//string com multiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Bruno`
	fmt.Println("O tamanho da string é", len(s2))

	//char??
	// var x char = 'b' - isso nao rola, o char não existe em go
	char := 'b'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
}
