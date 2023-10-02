package main

import "testing"

func TestSumarNumeros(t *testing.T)  {
	result := SumarNumeros(2,3)
	if result != 5 {
		t.Error("Resultado Incorrecto", result)
	}
}