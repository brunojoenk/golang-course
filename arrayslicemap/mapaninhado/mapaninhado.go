package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"J": {
			"João": 1000.00,
			"Jose": 2560.20,
		},
		"B": {
			"Bruno": 1200.00,
		},
	}

	delete(funcsPorLetra, "A")
	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
