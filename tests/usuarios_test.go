package tests

import "testing"

func TestBasico(t *testing.T) {
	if 2+2 != 4 {
		t.Error("Error en prueba")
	}
}
