package matematica

import "testing"

//Arquivos de teste devem terminar com _test.go
//Funções de teste devem iniciar com Test com o parâmetro (t *testing.T)

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

// go test cover - Mostra a cobertura de testes
// go test -coverprofile=resultado.out - Irá gerar um arquivo com os dados da cobertura de teste
// go tool cover -func=resultado.out - Mostra no terminal os dados de cobertura
// go tool cover -html=resultado.out - Mostra em uma pagina html o que foi coberto nos testes
