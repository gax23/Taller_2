package main

import "fmt"

type Alumno struct {
	Nombre    string
	Nota      int
	Siguiente *Alumno
}

func A(nombre string, nota int) *Alumno {
	return &Alumno{Nombre: nombre, Nota: nota}
}

func Desde(alumnos ...*Alumno) *Alumno {
	if len(alumnos) == 0 {
		return nil
	}
	for i := 0; i < len(alumnos)-1; i++ {
		alumnos[i].Siguiente = alumnos[i+1]
	}
	return alumnos[0]
}

func Str(a *Alumno) string {
	res := ""
	for n := a; n != nil; n = n.Siguiente {
		res += fmt.Sprintf("%s(%d) ", n.Nombre, n.Nota)
	}
	return res
}

func Partir(cabeza *Alumno) (*Alumno, *Alumno) {
	if cabeza == nil || cabeza.Siguiente == nil {
		return cabeza, nil
	}
	lento := cabeza
	rapido := cabeza.Siguiente
	for rapido != nil && rapido.Siguiente != nil {
		lento = lento.Siguiente
		rapido = rapido.Siguiente.Siguiente
	}
	mitad := lento.Siguiente
	lento.Siguiente = nil
	return cabeza, mitad
}

func Mezclar(a, b *Alumno) *Alumno {
	centinela := &Alumno{}
	actual := centinela
	for a != nil && b != nil {
		if a.Nota >= b.Nota {
			actual.Siguiente = a
			a = a.Siguiente
		} else {
			actual.Siguiente = b
			b = b.Siguiente
		}
		actual = actual.Siguiente
	}
	if a != nil {
		actual.Siguiente = a
	} else {
		actual.Siguiente = b
	}
	return centinela.Siguiente
}

func MergeSort(cabeza *Alumno) *Alumno {
	if cabeza == nil || cabeza.Siguiente == nil {
		return cabeza
	}
	izq, der := Partir(cabeza)
	izq = MergeSort(izq)
	der = MergeSort(der)
	return Mezclar(izq, der)
}

func InsertionSort(c *Alumno) *Alumno {
	if c == nil || c.Siguiente == nil {
		return c
	}
	ordenada := &Alumno{}
	actual := c
	for actual != nil {
		sig := actual.Siguiente
		ant := ordenada
		for ant.Siguiente != nil && ant.Siguiente.Nota >= actual.Nota {
			ant = ant.Siguiente
		}
		actual.Siguiente = ant.Siguiente
		ant.Siguiente = actual
		actual = sig
	}
	return ordenada.Siguiente
}

func main() {
	m := Desde(A("Ana", 15), A("Beto", 18), A("Caro", 15), A("Dani", 12))
	res := MergeSort(m)
	fmt.Println(Str(res))
}

/*
Ejercicio 6 - Parte B:
Hipótesis: La condición `a.Nota > b.Nota` en Mezclar hace que, ante notas iguales, se tome primero el elemento de la lista derecha (b), invirtiendo el orden original de inscripción.
Corrección mínima: cambiar `>` por `>=` para que, ante empate, se tome primero el elemento de la lista izquierda (a), preservando la estabilidad.
Test que rompe: TestMezclarEstable con notas empatadas. El original produce "Eva Beto Fito Caro Ana Dani"; el corregido "Beto Eva Ana Caro Fito Dani".

Parte C:
InsertionSort estable inserta cada nodo en la posición correcta manteniendo el orden relativo de los empates, usando `>=` para no adelantar elementos iguales.
*/
