package arquitetura

import (
	"runtime"
	"testing"
)

//go test std, roda teste dos pacotes da api de go

func TestDependente(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura adm64")
	}
	// ...
	t.Fail()
}

// pwd pasta de testes executar o comando abaixo para executar os testes daquele pacote
// go test ./...
