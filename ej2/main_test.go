package main

import "testing"

func TestEliminarTodos(t *testing.T) {
	l := Desde(1, 2, 3, 2, 4, 2)
	n := l.EliminarTodos(2)
	if n != 3 {
		t.Fatalf("borrados = %d", n)
	}
	if got := l.String(); got != "1 -> 3 -> 4 -> nil" {
		t.Fatalf("String() = %q", got)
	}
}

func TestEliminarTodosCabeza(t *testing.T) {
	l := Desde(5, 5, 5, 1, 2)
	n := l.EliminarTodos(5)
	if n != 3 {
		t.Fatalf("borrados = %d", n)
	}
	if got := l.String(); got != "1 -> 2 -> nil" {
		t.Fatalf("String() = %q", got)
	}
}

func TestEliminarTodosConsecutivos(t *testing.T) {
	l := Desde(7, 7, 2, 7, 7, 7, 4)
	n := l.EliminarTodos(7)
	if n != 5 {
		t.Fatalf("borrados = %d", n)
	}
	if got := l.String(); got != "2 -> 4 -> nil" {
		t.Fatalf("String() = %q", got)
	}
}

func TestEliminarDuplicados(t *testing.T) {
	l := Desde(1, 2, 1, 3, 2, 4, 1)
	l.EliminarDuplicados()
	if got := l.String(); got != "1 -> 2 -> 3 -> 4 -> nil" {
		t.Fatalf("String() = %q", got)
	}
	l2 := Desde(5, 5, 5, 5)
	l2.EliminarDuplicados()
	if got := l2.String(); got != "5 -> nil" {
		t.Fatalf("String() = %q", got)
	}
	l3 := Desde()
	l3.EliminarDuplicados()
	if got := l3.String(); got != "nil" {
		t.Fatalf("String() = %q", got)
	}
}

/*
Los tests cubren listas vacías, todos los elementos iguales, duplicados consecutivos y no consecutivos, asegurando que EliminarTodos borra todas las apariciones y que EliminarDuplicados conserva el orden.
*/
