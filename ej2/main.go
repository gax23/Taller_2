package main

import "fmt"

type Nodo struct {
	Valor     int
	Siguiente *Nodo
}

type Lista struct {
	Cabeza *Nodo
	n      int
}

func Desde(vals ...int) *Lista {
	l := &Lista{}
	for _, v := range vals {
		l.InsertarFinal(v)
	}
	return l
}

func (l *Lista) InsertarFinal(v int) {
	nuevo := &Nodo{Valor: v}
	if l.Cabeza == nil {
		l.Cabeza = nuevo
		l.n++
		return
	}
	actual := l.Cabeza
	for actual.Siguiente != nil {
		actual = actual.Siguiente
	}
	actual.Siguiente = nuevo
	l.n++
}

func (l *Lista) EliminarTodos(v int) int {
	borrados := 0
	for l.Cabeza != nil && l.Cabeza.Valor == v {
		l.Cabeza = l.Cabeza.Siguiente
		borrados++
	}
	if l.Cabeza == nil {
		return borrados
	}
	ant := l.Cabeza
	for ant.Siguiente != nil {
		if ant.Siguiente.Valor == v {
			ant.Siguiente = ant.Siguiente.Siguiente
			borrados++
		} else {
			ant = ant.Siguiente
		}
	}
	return borrados
}

func (l *Lista) String() string {
	if l.Cabeza == nil {
		return "nil"
	}
	res := ""
	for a := l.Cabeza; a != nil; a = a.Siguiente {
		res += fmt.Sprintf("%d -> ", a.Valor)
	}
	return res + "nil"
}

func (l *Lista) EliminarDuplicados() {
	actual := l.Cabeza
	for actual != nil {
		ant := actual
		for ant.Siguiente != nil {
			if ant.Siguiente.Valor == actual.Valor {
				ant.Siguiente = ant.Siguiente.Siguiente
			} else {
				ant = ant.Siguiente
			}
		}
		actual = actual.Siguiente
	}
}

func main() {
	l := Desde(1, 2, 1, 3, 2, 4)
	l.EliminarDuplicados()
	fmt.Println(l.String())
}

/*
Ejercicio 2 - Parte B:
Hipótesis: El defecto está en el avance de `ant` después de borrar. La línea `ant = ant.Siguiente` se ejecuta siempre, incluso cuando se acaba de eliminar un nodo. Si hay valores consecutivos iguales a v, el siguiente nodo (que también es v) se salta y no se elimina.
Corrección mínima: mover `ant = ant.Siguiente` a un `else`, de modo que solo avance cuando no se ha borrado.
Test que rompe: TestEliminarTodosConsecutivos con lista 7,7,2,7,7,7,4 espera 5 borrados y lista "2 -> 4 -> nil". El original devuelve 4 y deja un 7.

Parte C:
EliminarDuplicados() usa dos punteros: `actual` recorre la lista, `ant` recorre desde `actual` para borrar duplicados posteriores. Conserva la primera aparición.
*/
