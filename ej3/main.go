package main

import "fmt"

type NodoDoble struct {
	Valor     int
	Siguiente *NodoDoble
	Anterior  *NodoDoble
}

type ListaDoble struct {
	Cabeza *NodoDoble
	Cola   *NodoDoble
}

func (l *ListaDoble) InsertarOrdenado(v int) {
	nuevo := &NodoDoble{Valor: v}
	if l.Cabeza == nil {
		l.Cabeza = nuevo
		l.Cola = nuevo
		return
	}
	if v > l.Cabeza.Valor {
		nuevo.Siguiente = l.Cabeza
		l.Cabeza.Anterior = nuevo
		l.Cabeza = nuevo
		return
	}
	actual := l.Cabeza
	for actual.Siguiente != nil && actual.Siguiente.Valor >= v {
		actual = actual.Siguiente
	}
	nuevo.Siguiente = actual.Siguiente
	nuevo.Anterior = actual
	actual.Siguiente = nuevo
	if nuevo.Siguiente != nil {
		nuevo.Siguiente.Anterior = nuevo
	}
	if nuevo.Siguiente == nil {
		l.Cola = nuevo
	}
}

func (l *ListaDoble) Adelante() string {
	res := ""
	for a := l.Cabeza; a != nil; a = a.Siguiente {
		res += fmt.Sprintf("%d ", a.Valor)
	}
	return res
}

func (l *ListaDoble) Atras() string {
	res := ""
	for a := l.Cola; a != nil; a = a.Anterior {
		res += fmt.Sprintf("%d ", a.Valor)
	}
	return res
}

func (l *ListaDoble) Validar() error {
	if l.Cabeza == nil {
		if l.Cola != nil {
			return fmt.Errorf("cabeza nil pero cola no nil")
		}
		return nil
	}
	if l.Cola == nil {
		return fmt.Errorf("cabeza no nil pero cola nil")
	}
	var ultimo *NodoDoble
	for a := l.Cabeza; a != nil; a = a.Siguiente {
		if a.Siguiente != nil && a.Siguiente.Anterior != a {
			return fmt.Errorf("enlace anterior roto en %d", a.Valor)
		}
		ultimo = a
	}
	if ultimo != l.Cola {
		return fmt.Errorf("cola no es el ultimo nodo")
	}
	return nil
}

func (l *ListaDoble) EliminarNodo(n *NodoDoble) {
	if n == nil {
		return
	}
	if n.Anterior != nil {
		n.Anterior.Siguiente = n.Siguiente
	} else {
		l.Cabeza = n.Siguiente
	}
	if n.Siguiente != nil {
		n.Siguiente.Anterior = n.Anterior
	} else {
		l.Cola = n.Anterior
	}
	n.Anterior = nil
	n.Siguiente = nil
}

func main() {
	l := &ListaDoble{}
	for _, v := range []int{40, 10, 30, 20, 50} {
		l.InsertarOrdenado(v)
	}
	fmt.Println(l.Adelante())
	fmt.Println(l.Atras())
}

/*
Ejercicio 3 - Parte B:
Hipótesis: En InsertarOrdenado, al insertar un nodo intermedio, falta actualizar el puntero Anterior del nodo que queda después del nuevo. La línea `actual.Siguiente = nuevo` reconecta hacia adelante, pero `nuevo.Siguiente.Anterior` sigue apuntando al nodo anterior, rompiendo el recorrido hacia atrás.
Corrección mínima: agregar `if nuevo.Siguiente != nil { nuevo.Siguiente.Anterior = nuevo }`.
Test que rompe: TestInsertarOrdenadoAtras inserta 40,10,30,20,50. El original devuelve "50 40 10 " en Atras(); el corregido "50 40 30 20 10 ".

Parte C:
Validar() comprueba que para cada nodo, si tiene siguiente, su Anterior apunte a él, y que Cola sea el último nodo.
EliminarNodo() contempla cabeza, cola y nodo único, actualizando los cuatro punteros necesarios.
*/
