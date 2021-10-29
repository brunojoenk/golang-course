package main

import "fmt"

func notaPorConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough //executa os proximos cases, ou sejam não sai do switch
	case 9:
		return "A" //por padrão executa um case e sai do switch
		// nas outras linguagens precisa colocar break para parar, geralmente
		// saem executando todos os cases, o que não é o caso do golang
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaPorConceito(9.8))
	fmt.Println(notaPorConceito(6.8))
	fmt.Println(notaPorConceito(2.1))
	fmt.Println(notaPorConceito(-9.5))
}
