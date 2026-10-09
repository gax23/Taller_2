package main

import "testing"

func TestInsertarFinal(t *testing.T) {
	l := &Lista{}
	l.InsertarFinal(1)
	l.InsertarFinal(2)
	l.InsertarFinal(3)
	if got := l.String(); got != "1 -> 2 -> 3 -> nil" {
		t.Fatalf("String() = %q", got)
	}
	if l.Longitud() != 3 {
		t.Fatalf("Longitud() = %d", l.Longitud())
	}
}

func TestEliminarCabeza(t *testing.T) {
	l := &Lista{}
	l.InsertarFinal(10)
	l.InsertarFinal(20)
	l.Eliminar(10)
	if got := l.String(); got != "20 -> nil" {
		t.Fatalf("String() = %q", got)
	}
}

func TestEliminarCola(t *testing.T) {
	l := &Lista{}
	for _, v := range []int{10, 20, 30} {
		l.InsertarFinal(v)
	}
	l.Eliminar(30)
	l.InsertarFinal(40)
	if got := l.String(); got != "10 -> 20 -> 40 -> nil" {
		t.Fatalf("String() = %q", got)
	}
	if l.Longitud() != 3 {
		t.Fatalf("Longitud() = %d", l.Longitud())
	}
	if err := l.Validar(); err != nil {
		t.Fatal(err)
	}
}

func TestValidar(t *testing.T) {
	l := &Lista{}
	if err := l.Validar(); err != nil {
		t.Fatal(err)
	}
	l.InsertarFinal(1)
	if err := l.Validar(); err != nil {
		t.Fatal(err)
	}
	l.Cola = nil
	if err := l.Validar(); err == nil {
		t.Fatal("se esperaba error")
	}
}

func TestEliminarUltimo(t *testing.T) {
	l := &Lista{}
	l.InsertarFinal(5)
	l.InsertarFinal(6)
	v, err := l.EliminarUltimo()
	if err != nil || v != 6 {
		t.Fatalf("v=%d err=%v", v, err)
	}
	if got := l.String(); got != "5 -> nil" {
		t.Fatalf("String() = %q", got)
	}
	v, err = l.EliminarUltimo()
	if err != nil || v != 5 {
		t.Fatalf("v=%d err=%v", v, err)
	}
	if err := l.Validar(); err != nil {
		t.Fatal(err)
	}
	_, err = l.EliminarUltimo()
	if err == nil {
		t.Fatal("se esperaba error en lista vacia")
	}
}

/*
Los tests verifican el invariante de Cola tras eliminar el último nodo, la validación de consistencia y el correcto funcionamiento de EliminarUltimo en todos los casos.
*/
