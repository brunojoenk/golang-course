package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Jose":  1000.50,
		"Maria": 2100.75,
		"Ana":   3300.00, //precisa do virgula na ultima linha, ou fecha a chaves após o valor -> "Ana":   3300.00 }
	}

	funcsESalarios["Rafael"] = 1400.35
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
