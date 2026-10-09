package main

import "testing"

func TestInvertirEnGrupos(t *testing.T) {
	n := Desde(1, 2, 3, 4, 5, 6)
	res := InvertirEnGrupos(n, 2)
	if got := Str(res); got != "2 1 4 3 6 5 " {
		t.Fatalf("Str() = %q", got)
	}
}

func TestInvertirEnGruposIncompleto(t *testing.T) {
	n := Desde(1, 2, 3, 4, 5, 6, 7, 8)
	res := InvertirEnGrupos(n, 3)
	if got := Str(res); got != "3 2 1 6 5 4 7 8 " {
		t.Fatalf("Str() = %q", got)
	}
}

func TestEsPalindromo(t *testing.T) {
	n := Desde(1, 2, 3, 2, 1)
	if !EsPalindromo(n) {
		t.Fatal("deberia ser palindromo")
	}
	if got := Str(n); got != "1 2 3 2 1 " {
		t.Fatalf("lista modificada: %q", got)
	}
	n2 := Desde(1, 2, 3, 4)
	if EsPalindromo(n2) {
		t.Fatal("no deberia ser palindromo")
	}
	if got := Str(n2); got != "1 2 3 4 " {
		t.Fatalf("lista modificada: %q", got)
	}
}

/*
Los tests comprueban que los bloques incompletos no se invierten y que EsPalindromo funciona sin modificar la lista original.
*/
