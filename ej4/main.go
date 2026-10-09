package main

import "fmt"

type Nodo struct {
	Valor     int
	Siguiente *Nodo
}

func Desde(vals ...int) *Nodo {
	if len(vals) == 0 {
		return nil
	}
	cabeza := &Nodo{Valor: vals[0]}
	actual := cabeza
	for _, v := range vals[1:] {
		actual.Siguiente = &Nodo{Valor: v}
		actual = actual.Siguiente
	}
	return cabeza
}

func Str(n *Nodo) string {
	res := ""
	for a := n; a != nil; a = a.Siguiente {
		res += fmt.Sprintf("%d ", a.Valor)
	}
	return res
}

func InvertirEnGrupos(cabeza *Nodo, k int) *Nodo {
	if cabeza == nil || k <= 1 {
		return cabeza
	}
	var resultado *Nodo
	var ultimoGrupo *Nodo
	actual := cabeza
	for actual != nil {
		cont := 0
		tmp := actual
		for tmp != nil && cont < k {
			tmp = tmp.Siguiente
			cont++
		}
		if cont < k {
			break
		}
		inicioGrupo := actual
		var ant *Nodo
		cont = 0
		for actual != nil && cont < k {
			sig := actual.Siguiente
			actual.Siguiente = ant
			ant = actual
			actual = sig
			cont++
		}
		if resultado == nil {
			resultado = ant
		} else {
			ultimoGrupo.Siguiente = ant
		}
		ultimoGrupo = inicioGrupo
	}
	if ultimoGrupo != nil {
		ultimoGrupo.Siguiente = actual
	} else {
		resultado = actual
	}
	return resultado
}

func EsPalindromo(cabeza *Nodo) bool {
	if cabeza == nil || cabeza.Siguiente == nil {
		return true
	}
	lento, rapido := cabeza, cabeza
	for rapido.Siguiente != nil && rapido.Siguiente.Siguiente != nil {
		lento = lento.Siguiente
		rapido = rapido.Siguiente.Siguiente
	}
	var ant *Nodo
	actual := lento.Siguiente
	for actual != nil {
		sig := actual.Siguiente
		actual.Siguiente = ant
		ant = actual
		actual = sig
	}
	p1 := cabeza
	p2 := ant
	esPal := true
	for p2 != nil {
		if p1.Valor != p2.Valor {
			esPal = false
			break
		}
		p1 = p1.Siguiente
		p2 = p2.Siguiente
	}
	var prev *Nodo
	actual = ant
	for actual != nil {
		sig := actual.Siguiente
		actual.Siguiente = prev
		prev = actual
		actual = sig
	}
	lento.Siguiente = prev
	return esPal
}

func main() {
	n := Desde(1, 2, 3, 4, 5, 6, 7, 8)
	res := InvertirEnGrupos(n, 3)
	fmt.Println(Str(res))
}

/*
Ejercicio 4 - Parte B:
Hipótesis: El código invierte el bloque final aunque tenga menos de k nodos. El bucle interior termina por `actual == nil`, no por `cont == k`, y no se verifica el tamaño del bloque.
Corrección mínima: antes de invertir, contar hasta k nodos; si no se alcanzan, salir del bucle exterior sin invertir.
Test que rompe: TestInvertirEnGruposIncompleto con n=8, k=3. El original devuelve "3 2 1 6 5 4 8 7 "; el corregido "3 2 1 6 5 4 7 8 ".

Parte C:
EsPalindromo() encuentra el medio con punteros lento/rápido, invierte la segunda mitad, compara y luego la vuelve a invertir para dejar la lista intacta.
*/
