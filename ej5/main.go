package main

import "fmt"

type Proceso struct {
	Nombre    string
	Tiempo    int
	Llegada   int
	Rafaga    int
	Siguiente *Proceso
}

type Planificador struct {
	Cola       *Proceso
	Tamanio    int
	pendientes []*Proceso
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
	if p.Cola == nil {
		return nil
	}
	var orden []string
	reloj := 0
	ant := p.Cola
	act := p.Cola.Siguiente
	for p.Tamanio > 0 {
		usa := quantum
		if act.Tiempo < quantum {
			usa = act.Tiempo
		}
		act.Tiempo -= usa
		reloj += usa
		if act.Tiempo == 0 {
			orden = append(orden, fmt.Sprintf("%s@t=%d", act.Nombre, reloj))
			ant.Siguiente = act.Siguiente
			p.Tamanio--
		} else {
			ant = act
		}
		act = act.Siguiente
	}
	p.Cola = nil
	return orden
}

func (p *Planificador) insertarEnCola(nuevo *Proceso) {
	if p.Cola == nil {
		nuevo.Siguiente = nuevo
	} else {
		nuevo.Siguiente = p.Cola.Siguiente
		p.Cola.Siguiente = nuevo
	}
	p.Cola = nuevo
	p.Tamanio++
}

func (p *Planificador) AgregarConLlegada(nombre string, llegada, rafaga int) {
	nuevo := &Proceso{Nombre: nombre, Tiempo: rafaga, Llegada: llegada, Rafaga: rafaga}
	p.pendientes = append(p.pendientes, nuevo)
}

func (p *Planificador) EjecutarConLlegadas(quantum int) ([]string, float64) {
	var orden []string
	reloj := 0
	esperaTotal := 0
	terminados := 0
	for i := 0; i < len(p.pendientes); i++ {
		for j := i + 1; j < len(p.pendientes); j++ {
			if p.pendientes[j].Llegada < p.pendientes[i].Llegada {
				p.pendientes[i], p.pendientes[j] = p.pendientes[j], p.pendientes[i]
			}
		}
	}
	idx := 0
	for idx < len(p.pendientes) && p.pendientes[idx].Llegada <= reloj {
		p.insertarEnCola(p.pendientes[idx])
		idx++
	}
	if p.Tamanio == 0 && idx < len(p.pendientes) {
		reloj = p.pendientes[idx].Llegada
		for idx < len(p.pendientes) && p.pendientes[idx].Llegada <= reloj {
			p.insertarEnCola(p.pendientes[idx])
			idx++
		}
	}
	ant := p.Cola
	act := p.Cola.Siguiente
	for p.Tamanio > 0 || idx < len(p.pendientes) {
		if p.Tamanio == 0 {
			reloj = p.pendientes[idx].Llegada
			for idx < len(p.pendientes) && p.pendientes[idx].Llegada <= reloj {
				p.insertarEnCola(p.pendientes[idx])
				idx++
			}
			ant = p.Cola
			act = p.Cola.Siguiente
			continue
		}
		usa := quantum
		if act.Tiempo < quantum {
			usa = act.Tiempo
		}
		act.Tiempo -= usa
		reloj += usa
		for idx < len(p.pendientes) && p.pendientes[idx].Llegada <= reloj {
			p.insertarEnCola(p.pendientes[idx])
			idx++
		}
		if act.Tiempo == 0 {
			orden = append(orden, fmt.Sprintf("%s@t=%d", act.Nombre, reloj))
			esperaTotal += reloj - act.Llegada - act.Rafaga
			terminados++
			ant.Siguiente = act.Siguiente
			p.Tamanio--
		} else {
			ant = act
		}
		act = act.Siguiente
	}
	p.Cola = nil
	p.pendientes = nil
	var promedio float64
	if terminados > 0 {
		promedio = float64(esperaTotal) / float64(terminados)
	}
	return orden, promedio
}

func main() {
	b := &Planificador{}
	b.Agregar("P1", 2)
	b.Agregar("P2", 2)
	b.Agregar("P3", 5)
	res := b.Ejecutar(2)
	for _, s := range res {
		fmt.Println(s)
	}
}

/*
Ejercicio 5 - Parte B:
Hipótesis: Al retirar un proceso, `ant` se actualiza a `act` (el nodo eliminado) y luego `act` avanza. En la siguiente eliminación, `ant.Siguiente` ya no pertenece al anillo, por lo que el proceso terminado no se desenlaza realmente y vuelve a ejecutarse.
Corrección mínima: solo actualizar `ant = act` cuando el proceso no termina (else). Así `ant` permanece como el nodo anterior al siguiente.
Test que rompe: TestEjecutarCorregido con P1=2, P2=2, P3=5, quantum=2. El original devuelve [P1@t=2 P2@t=4 P2@t=6]; el corregido [P1@t=2 P2@t=4 P3@t=9].

Parte C:
Se añade Llegada y Rafaga a Proceso. AgregarConLlegada almacena en pendientes. EjecutarConLlegadas ordena por llegada, inserta cuando reloj >= llegada y calcula espera promedio.
*/
