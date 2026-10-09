# Informe del Taller número 2

**Integrantes:** Gael Almeyda, Alvaro Saavedra
**Profesor:** Marks
**Sección:** S-002

Este informe de predicciones corresponde a la parte A de la actividad "Taller 2" y fue elaborado antes de ejecutar el código. Nota: las salidas de los ejercicios 3, 4 y 6 terminan cada línea con un espacio, porque así construyen sus cadenas las funciones `Adelante`, `Atras` y `Str`.

---

## Ejercicio 1: Lista simple con puntero Cola

**Código:**

```go
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
	if l.Cabeza == nil {
		l.Cabeza, l.Cola = nuevo, nuevo
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
			l.n--
			return true
		}
		ant = ant.Siguiente
	}
	return false
}

func (l *Lista) Longitud() int { return l.n }

func (l *Lista) String() string {
	s := ""
	for a := l.Cabeza; a != nil; a = a.Siguiente {
		s += fmt.Sprintf("%d -> ", a.Valor)
	}
	return s + "nil"
}

func main() {
	l := &Lista{}
	for _, v := range []int{10, 20, 30, 40} {
		l.InsertarFinal(v)
	}
	l.Eliminar(20)
	l.InsertarFinal(50)
	fmt.Println(l, "│ n =", l.Longitud())
	l.Eliminar(50)
	l.InsertarFinal(60)
	fmt.Println(l, "│ n =", l.Longitud())
}
```

**Predicción:**

```
10 -> 30 -> 40 -> 50 -> nil │ n = 4
10 -> 30 -> 40 -> nil │ n = 4
```

**Análisis:**

- Tras insertar 10, 20, 30, 40 y eliminar 20 (nodo intermedio), la lista queda `10 → 30 → 40` con n = 3. Al insertar 50 queda `10 → 30 → 40 → 50` con n = 4, lo que da la primera línea.
- `Eliminar(50)` elimina el último nodo por la rama general y solo reajusta `ant.Siguiente`, sin tocar `Cola`. Por eso `Cola` sigue apuntando al nodo 50, que ya no está en la lista.
- `InsertarFinal(60)` cuelga el 60 de ese nodo desconectado. Desde `Cabeza` el 60 no se alcanza, así que la lista impresa no lo muestra, pero `n` sí sube a 4.

---

## Ejercicio 2: Eliminar todas las apariciones

**Código:**

```go
package main

import "fmt"

type Nodo struct {
	Valor     int
	Siguiente *Nodo
}

type Lista struct{ Cabeza *Nodo }

func Desde(vals ...int) *Lista {
	l := &Lista{}
	var cola *Nodo
	for _, v := range vals {
		n := &Nodo{Valor: v}
		if cola == nil {
			l.Cabeza = n
		} else {
			cola.Siguiente = n
		}
		cola = n
	}
	return l
}

func (l *Lista) String() string {
	s := ""
	for a := l.Cabeza; a != nil; a = a.Siguiente {
		s += fmt.Sprintf("%d -> ", a.Valor)
	}
	return s + "nil"
}

func (l *Lista) EliminarTodos(v int) int {
	borrados := 0
	for l.Cabeza != nil && l.Cabeza.Valor == v {
		l.Cabeza = l.Cabeza.Siguiente
		borrados++
	}
	ant := l.Cabeza
	for ant != nil && ant.Siguiente != nil {
		if ant.Siguiente.Valor == v {
			ant.Siguiente = ant.Siguiente.Siguiente
			borrados++
		}
		ant = ant.Siguiente
	}
	return borrados
}

func main() {
	a := Desde(3, 7, 1, 7, 5, 7)
	fmt.Println(a.EliminarTodos(7), a)
	b := Desde(7, 7, 2, 7, 7, 7, 4)
	fmt.Println(b.EliminarTodos(7), b)
}
```

**Predicción:**

```
3 3 -> 1 -> 5 -> nil
4 2 -> 7 -> 4 -> nil
```

**Análisis:**

- **Caso a:** no hay 7 consecutivos. Cada vez que se borra un 7, `ant` avanza al nodo siguiente, que nunca es otro 7, así que se borran los tres y queda `3 → 1 → 5`.
- **Caso b:** el primer bucle quita los dos 7 iniciales y deja `2 → 7 → 7 → 7 → 4`. En el segundo bucle, `ant` avanza siempre, incluso después de borrar. Se borra un 7, `ant` pasa al siguiente 7, se borra otro, y `ant` salta al 4 sin revisar el último 7. En total se borran 4 (2 + 2), no 5, y queda un 7 en la lista.

---

## Ejercicio 3: Lista doble ordenada

**Código:**

```go
package main

import "fmt"

type NodoDoble struct {
	Valor     int
	Anterior  *NodoDoble
	Siguiente *NodoDoble
}

type ListaDoble struct {
	Cabeza *NodoDoble
	Cola   *NodoDoble
}

func (l *ListaDoble) InsertarOrdenado(v int) {
	nuevo := &NodoDoble{Valor: v}
	if l.Cabeza == nil {
		l.Cabeza, l.Cola = nuevo, nuevo
		return
	}
	if v < l.Cabeza.Valor {
		nuevo.Siguiente = l.Cabeza
		l.Cabeza.Anterior = nuevo
		l.Cabeza = nuevo
		return
	}
	actual := l.Cabeza
	for actual.Siguiente != nil && actual.Siguiente.Valor < v {
		actual = actual.Siguiente
	}
	nuevo.Siguiente = actual.Siguiente
	nuevo.Anterior = actual
	actual.Siguiente = nuevo
	if nuevo.Siguiente == nil {
		l.Cola = nuevo
	}
}

func (l *ListaDoble) Adelante() string {
	s := ""
	for a := l.Cabeza; a != nil; a = a.Siguiente {
		s += fmt.Sprintf("%d ", a.Valor)
	}
	return s
}

func (l *ListaDoble) Atras() string {
	s := ""
	for a := l.Cola; a != nil; a = a.Anterior {
		s += fmt.Sprintf("%d ", a.Valor)
	}
	return s
}

func main() {
	l := &ListaDoble{}
	for _, v := range []int{10, 20, 30, 40, 50} {
		l.InsertarOrdenado(v)
	}
	fmt.Println("A:", l.Adelante(), "│ B:", l.Atras())
	m := &ListaDoble{}
	for _, v := range []int{40, 10, 30, 20, 50} {
		m.InsertarOrdenado(v)
	}
	fmt.Println("A:", m.Adelante(), "│ B:", m.Atras())
}
```

**Predicción:**

```
A: 10 20 30 40 50  │ B: 50 40 30 20 10 
A: 10 20 30 40 50  │ B: 50 40 10 
```

**Análisis:**

- **Caso l:** los valores llegan ya ordenados, así que todos entran al final de la lista. Esa rama sí actualiza `Cola`, y el recorrido hacia atrás sale completo.
- **Caso m:** el 10 entra por la rama de la cabeza. Luego el 30 y el 20 entran como nodos intermedios. La rama intermedia no actualiza `actual.Siguiente.Anterior`, así que el nodo 40 conserva `Anterior = 10`, y el 30 conserva `Anterior = 10`.
- El recorrido hacia adelante sale bien (`10 20 30 40 50`), pero hacia atrás se salta el 30 y el 20: `50 → 40 → 10`.
- Las dos líneas tienen doble espacio antes de `│` porque `Adelante()` ya termina en espacio y `Println` agrega otro.

---

## Ejercicio 4: Invertir en grupos de k

**Código:**

```go
package main

import "fmt"

type Nodo struct {
	Valor     int
	Siguiente *Nodo
}

func Desde(vals ...int) *Nodo {
	var cab, cola *Nodo
	for _, v := range vals {
		n := &Nodo{Valor: v}
		if cola == nil {
			cab = n
		} else {
			cola.Siguiente = n
		}
		cola = n
	}
	return cab
}

func Str(c *Nodo) string {
	s := ""
	for ; c != nil; c = c.Siguiente {
		s += fmt.Sprintf("%d ", c.Valor)
	}
	return s
}

func InvertirEnGrupos(cabeza *Nodo, k int) *Nodo {
	centinela := &Nodo{Siguiente: cabeza}
	colaPrevia := centinela
	actual := cabeza
	for actual != nil {
		inicioGrupo := actual
		var ant *Nodo
		cont := 0
		for actual != nil && cont < k {
			sig := actual.Siguiente
			actual.Siguiente = ant
			ant = actual
			actual = sig
			cont++
		}
		colaPrevia.Siguiente = ant
		inicioGrupo.Siguiente = actual
		colaPrevia = inicioGrupo
	}
	return centinela.Siguiente
}

func main() {
	fmt.Println(Str(InvertirEnGrupos(Desde(1, 2, 3, 4, 5, 6), 3)))
	fmt.Println(Str(InvertirEnGrupos(Desde(1, 2, 3, 4, 5, 6, 7, 8), 3)))
}
```

**Predicción:**

```
3 2 1 6 5 4 
3 2 1 6 5 4 8 7 
```

**Análisis:**

- **n = 6, k = 3:** hay dos bloques completos, `[1 2 3]` y `[4 5 6]`. Cada uno se invierte y se reconecta con el anterior: `3 2 1 6 5 4`.
- **n = 8, k = 3:** los dos primeros bloques salen igual. Sobran dos nodos (`7 8`), pero el bucle interior termina por `actual == nil`, no por `cont == k`, y el código no revisa si el bloque quedó incompleto. Invierte igual y da `8 7`. Según la especificación debería quedar `7 8`.

---

## Ejercicio 5: Planificador Round Robin

**Código:**

```go
package main

import "fmt"

type Proceso struct {
	Nombre    string
	Tiempo    int 
	Siguiente *Proceso
}

type Planificador struct {
	Cola    *Proceso
	Tamanio int
}

func (p *Planificador) Agregar(nombre string, tiempo int) {
	nuevo := &Proceso{Nombre: nombre, Tiempo: tiempo}
	if p.Cola == nil {
		nuevo.Siguiente = nuevo
	} else {
		nuevo.Siguiente = p.Cola.Siguiente
		p.Cola.Siguiente = nuevo
	}
	p.Cola = nuevo
	p.Tamanio++
}


func (p *Planificador) Ejecutar(quantum int) []string {
	var orden []string
	reloj := 0
	ant := p.Cola
	act := p.Cola.Siguiente
	for p.Tamanio > 0 {
		usa := min(quantum, act.Tiempo)
		act.Tiempo -= usa
		reloj += usa
		if act.Tiempo == 0 {
			orden = append(orden, fmt.Sprintf("%s@t=%d", act.Nombre, reloj))
			ant.Siguiente = act.Siguiente 
			p.Tamanio--
		}
		ant = act
		act = act.Siguiente
	}
	p.Cola = nil
	return orden
}

func main() {
	a := &Planificador{}
	a.Agregar("P1", 6)
	a.Agregar("P2", 4)
	a.Agregar("P3", 8)
	fmt.Println(a.Ejecutar(2))

	b := &Planificador{}
	b.Agregar("P1", 2)
	b.Agregar("P2", 2)
	b.Agregar("P3", 5)
	fmt.Println(b.Ejecutar(2))
}
```

**Predicción:**

```
[P2@t=10 P1@t=14 P3@t=18]
[P1@t=2 P2@t=4 P2@t=6]
```

**Análisis:**

- **Caso a (quantum = 2):** el orden de turnos es P1, P2, P3, P1, P2. P2 termina en t = 10 y se desenlaza bien, porque el siguiente en terminar es P1 y su `ant` es válido. P1 termina en t = 14 y P3 en t = 18, que coincide con la suma de ráfagas (6 + 4 + 8).
- **Caso b:** P1 termina en t = 2 y P2 en t = 4. Tras retirar P1, `ant` queda apuntando a P1, que ya no está en el anillo. Cuando P2 termina, se ejecuta `P1.Siguiente = P2.Siguiente`, que no desenlaza a P2 de verdad, porque `P3.Siguiente` sigue apuntando a P2.
- P3 corre un turno (t = 6, le quedan 3), el recorrido vuelve a caer en P2, que ya terminó, y se registra otra vez como `P2@t=6`. Con eso `Tamanio` llega a 0 y el ciclo se detiene: P3 nunca termina.

---

## Ejercicio 6: Merge Sort estable en lista

**Código:**

```go
package main

import "fmt"

type Alumno struct {
	Nombre    string
	Nota      int
	Siguiente *Alumno
}

func Desde(datos ...Alumno) *Alumno {
	var cab, cola *Alumno
	for i := range datos {
		n := &Alumno{Nombre: datos[i].Nombre, Nota: datos[i].Nota}
		if cola == nil {
			cab = n
		} else {
			cola.Siguiente = n
		}
		cola = n
	}
	return cab
}

func A(nombre string, nota int) Alumno { return Alumno{Nombre: nombre, Nota: nota} }

func Str(c *Alumno) string {
	s := ""
	for ; c != nil; c = c.Siguiente {
		s += fmt.Sprintf("%s(%d) ", c.Nombre, c.Nota)
	}
	return s
}

func ObtenerMedio(c *Alumno) *Alumno {
	lento, rapido := c, c.Siguiente
	for rapido != nil && rapido.Siguiente != nil {
		lento = lento.Siguiente
		rapido = rapido.Siguiente.Siguiente
	}
	return lento
}


func Mezclar(a, b *Alumno) *Alumno {
	centinela := &Alumno{}
	cola := centinela
	for a != nil && b != nil {
		if a.Nota > b.Nota {
			cola.Siguiente, a = a, a.Siguiente
		} else {
			cola.Siguiente, b = b, b.Siguiente
		}
		cola = cola.Siguiente
	}
	if a != nil {
		cola.Siguiente = a
	} else {
		cola.Siguiente = b
	}
	return centinela.Siguiente
}


func MergeSort(c *Alumno) *Alumno {
	if c == nil || c.Siguiente == nil {
		return c
	}
	medio := ObtenerMedio(c)
	der := medio.Siguiente
	medio.Siguiente = nil
	return Mezclar(MergeSort(c), MergeSort(der))
}

func main() {
	l := Desde(A("Ana", 14), A("Beto", 18), A("Caro", 11), A("Dani", 16))
	fmt.Println(Str(MergeSort(l)))
	m := Desde(A("Ana", 15), A("Beto", 18), A("Caro", 15), A("Dani", 12), A("Eva", 18), A("Fito", 15))
	fmt.Println(Str(MergeSort(m)))
}
```

**Predicción:**

```
Beto(18) Dani(16) Ana(14) Caro(11) 
Eva(18) Beto(18) Fito(15) Caro(15) Ana(15) Dani(12) 
```

**Análisis:**

- **Caso l:** todas las notas son distintas, así que no hay empates y el orden descendente sale correcto.
- **Caso m:** hay empates en 18 (Beto, Eva) y en 15 (Ana, Caro, Fito). En `Mezclar` la condición es `a.Nota > b.Nota`. Con notas iguales cae al `else` y toma el elemento de la lista derecha `b` antes que el de la izquierda `a`, lo que invierte el orden original de los empates.
- En la mezcla final, Eva (derecha) sale antes que Beto (izquierda), y dentro de las sublistas Caro sale antes que Ana. El resultado es `Eva Beto Fito Caro Ana Dani`, en lugar del estable `Beto Eva Ana Caro Fito Dani`.
- El enunciado del taller muestra `││` en `c == nil ││ c.Siguiente == nil`. Lo tomé como un error de formato del documento y lo interpreté como el operador `||`.
