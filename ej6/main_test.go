package main

import "testing"

func TestMergeSort(t *testing.T) {
	m := Desde(A("Ana", 15), A("Beto", 18), A("Caro", 10))
	res := MergeSort(m)
	if got := Str(res); got != "Beto(18) Ana(15) Caro(10) " {
		t.Fatalf("Str() = %q", got)
	}
}

func TestMezclarEstable(t *testing.T) {
	m := Desde(A("Ana", 15), A("Beto", 18), A("Caro", 15), A("Dani", 12), A("Eva", 18), A("Fito", 15))
	res := MergeSort(m)
	esperado := "Beto(18) Eva(18) Ana(15) Caro(15) Fito(15) Dani(12) "
	if got := Str(res); got != esperado {
		t.Fatalf("Str() = %q", got)
	}
}

func TestInsertionSortEstable(t *testing.T) {
	m := Desde(A("Ana", 15), A("Beto", 18), A("Caro", 15), A("Dani", 12), A("Eva", 18), A("Fito", 15))
	res := InsertionSort(m)
	esperado := "Beto(18) Eva(18) Ana(15) Caro(15) Fito(15) Dani(12) "
	if got := Str(res); got != esperado {
		t.Fatalf("Str() = %q", got)
	}
}

/*
Los tests verifican que MergeSort e InsertionSort producen exactamente el mismo orden estable ante empates.
*/
