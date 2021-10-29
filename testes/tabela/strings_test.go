package strings

//go test
//go test -v modo verboso, mostrar detalhes de log por exemplo

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)"

func TestIndex(t *testing.T) {
	// t.Parallel() executar os testes paralelamente, go routines vai fazer o paralelismo
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Cod3r é show", "Cod3r", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"bruno", "u", 2},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
