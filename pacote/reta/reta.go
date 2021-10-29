package main

import "math"

// Iniciando com letra maiúscula é PUBLICO, ou seja, vísivel dentro e fora do pacote
// Não pode ter funções com nomes iguais dentro dos pacotes
// Visibilidade privada dentro do pacote

// Iniciando com letra minuscula é PRIVADO, visivel apenas dentro do pacote
// Quando é compilado, todos os arquivos viram um pacote

// Ponto - gerará algo publico
// ponto ou _Ponto - gerará algo privado

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é responsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
