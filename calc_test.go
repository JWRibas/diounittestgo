package main

import "testing"

func TestSoma(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 6

	if teste != resultado {
		t.Error("Valor Esperado: ", resultado, "Valor retornado: ", teste)
	}
}
func TestMultiplica(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 100

	if teste != resultado {
		t.Error("Valor Esperado: ", resultado, "Valor retornado: ", teste)
	}
}
func TestDivide(t *testing.T) {
	teste := divide(15, 15)
	resultado := 1

	if teste != resultado {
		t.Error("Valor Esperado: ", resultado, "Valor retornado: ", teste)
	}
}
func TestSubtrai(t *testing.T) {
	teste := subtrai(20, 3)
	resultado := 17

	if teste != resultado {
		t.Error("Valor Esperado: ", resultado, "Valor retornado: ", teste)
	}
}
