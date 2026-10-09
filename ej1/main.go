package main

import "fmt"

type Nodo struct {
	Valor     int
	Siguiente *Nodo
}

type Lista struct {
	Cabeza *Nodo
	Cola   *Nodo
	n      int
}

func (l *Lista) InsertarFinal(v int) {
	nuevo := &Nodo{Valor: v}
	if l.Cola == nil {
		l.Cabeza = nuevo
		l.Cola = nuevo
	} else {
		l.Cola.Siguiente = nuevo
		l.Cola = nuevo
	}
	l.n++
}

func (l *Lista) Eliminar(v int) bool {
	if l.Cabeza == nil {
		return false
	}
	if l.Cabeza.Valor == v {
		l.Cabeza = l.Cabeza.Siguiente
		if l.Cabeza == nil {
			l.Cola = nil
		}
		l.n--
		return true
	}
	ant := l.Cabeza
	for ant.Siguiente != nil {
		if ant.Siguiente.Valor == v {
			ant.Siguiente = ant.Siguiente.Siguiente
			if ant.Siguiente == nil {
				l.Cola = ant
			}
			l.n--
			return true
		}
		ant = ant.Siguiente
	}
	return false
}

func (l *Lista) String() string {
	res := ""
	for a := l.Cabeza; a != nil; a = a.Siguiente {
		res += fmt.Sprintf("%d -> ", a.Valor)
	}
	return res + "nil"
}

func (l *Lista) Longitud() int {
	return l.n
}

func (l *Lista) Validar() error {
	if l.Cabeza == nil {
		if l.Cola != nil {
			return fmt.Errorf("cabeza nil pero cola no nil")
		}
		if l.n != 0 {
			return fmt.Errorf("cabeza nil pero n=%d", l.n)
		}
		return nil
	}
	if l.Cola == nil {
		return fmt.Errorf("cabeza no nil pero cola nil")
	}
	cont := 0
	var ultimo *Nodo
	for a := l.Cabeza; a != nil; a = a.Siguiente {
		cont++
		ultimo = a
	}
	if cont != l.n {
		return fmt.Errorf("n=%d pero nodos=%d", l.n, cont)
	}
	if ultimo != l.Cola {
		return fmt.Errorf("cola no es el ultimo nodo")
	}
	return nil
}

func (l *Lista) EliminarUltimo() (int, error) {
	if l.Cabeza == nil {
		return 0, fmt.Errorf("lista vacia")
	}
	if l.Cabeza == l.Cola {
		v := l.Cabeza.Valor
		l.Cabeza = nil
		l.Cola = nil
		l.n = 0
		return v, nil
	}
	ant := l.Cabeza
	for ant.Siguiente != l.Cola {
		ant = ant.Siguiente
	}
	v := l.Cola.Valor
	ant.Siguiente = nil
	l.Cola = ant
	l.n--
	return v, nil
}

func main() {
	l := &Lista{}
	l.InsertarFinal(10)
	l.InsertarFinal(20)
	l.InsertarFinal(30)
	fmt.Println(l.String())
}

/*
Ejercicio 1 - Parte B:
Hipótesis: El defecto está en Eliminar cuando se borra el último nodo por la rama general. La línea `ant.Siguiente = ant.Siguiente.Siguiente` deja a `l.Cola` apuntando al nodo eliminado. La corrección mínima añade `if ant.Siguiente == nil { l.Cola = ant }` justo después de reasignar.
Test que rompe: TestEliminarCola inserta 10,20,30, elimina 30, inserta 40. Con el código original, String() muestra "10 -> 20 -> nil" y Longitud() es 3, pero el 40 no es alcanzable. Con la corrección, muestra "10 -> 20 -> 40 -> nil".

Parte C:
Validar() recorre la lista contando nodos y verificando que el último coincida con Cola, que n coincida con el conteo y que Cabeza==nil ⇔ Cola==nil.
EliminarUltimo() maneja lista vacía, un solo nodo y varios nodos. Para varios, recorre hasta el penúltimo y actualiza Cola.
*/
