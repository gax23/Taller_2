package main

import "testing"

func TestInsertarOrdenadoAdelante(t *testing.T) {
	l := &ListaDoble{}
	for _, v := range []int{40, 10, 30, 20, 50} {
		l.InsertarOrdenado(v)
	}
	if got := l.Adelante(); got != "50 40 30 20 10 " {
		t.Fatalf("Adelante() = %q", got)
	}
}

func TestInsertarOrdenadoAtras(t *testing.T) {
	l := &ListaDoble{}
	for _, v := range []int{40, 10, 30, 20, 50} {
		l.InsertarOrdenado(v)
	}
	if got := l.Atras(); got != "10 20 30 40 50 " {
		t.Fatalf("Atras() = %q", got)
	}
	if err := l.Validar(); err != nil {
		t.Fatal(err)
	}
}

func TestEliminarNodo(t *testing.T) {
	l := &ListaDoble{}
	for _, v := range []int{10, 20, 30} {
		l.InsertarOrdenado(v)
	}
	l.EliminarNodo(l.Cabeza)
	if got := l.Adelante(); got != "20 10 " {
		t.Fatalf("Adelante() = %q", got)
	}
	if err := l.Validar(); err != nil {
		t.Fatal(err)
	}
	l.EliminarNodo(l.Cola)
	if got := l.Adelante(); got != "20 " {
		t.Fatalf("Adelante() = %q", got)
	}
	if err := l.Validar(); err != nil {
		t.Fatal(err)
	}
	l.EliminarNodo(l.Cabeza)
	if l.Cabeza != nil || l.Cola != nil {
		t.Fatal("lista no quedo vacia")
	}
}

/*
Los tests verifican el recorrido hacia atrás tras inserciones intermedias y la eliminación de nodos en cabeza, cola y único.
*/
