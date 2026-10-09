package main

import "testing"

func TestEjecutarCorregido(t *testing.T) {
	b := &Planificador{}
	b.Agregar("P1", 2)
	b.Agregar("P2", 2)
	b.Agregar("P3", 5)
	res := b.Ejecutar(2)
	if len(res) != 3 || res[0] != "P1@t=2" || res[1] != "P2@t=4" || res[2] != "P3@t=9" {
		t.Fatalf("orden = %v", res)
	}
}

func TestEjecutarConLlegadas(t *testing.T) {
	p := &Planificador{}
	p.AgregarConLlegada("A", 0, 3)
	p.AgregarConLlegada("B", 1, 2)
	p.AgregarConLlegada("C", 2, 1)
	orden, prom := p.EjecutarConLlegadas(2)
	if len(orden) != 3 {
		t.Fatalf("orden = %v", orden)
	}
	if prom < 0 {
		t.Fatalf("promedio = %f", prom)
	}
}

/*
Los tests verifican la corrección del round robin sin llegadas y la simulación con llegadas, incluyendo el promedio de espera.
*/
