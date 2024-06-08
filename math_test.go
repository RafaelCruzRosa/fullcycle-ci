package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(1, 1)

	if total != 2 {
		t.Errorf("Valor esperado %d e recebeu: %d", 2, total)
	}
}
