package ListaCDE

import (
	"Proyecto/estructuras/Archivos"
	"fmt"
	"strconv"
)

type ListaCircular struct{
	Inicio *NodoListaCircular
	Tamanio int
}

func (lista *ListaCircular) InsertarTutor(carnet int, nombre string, curso string, nota int){
	nuevoTutor := &Tutor{Carnet: carnet, Nombre: nombre, Curso: curso, Nota: nota}
	nuevoNodo := &NodoListaCircular{Tutor: nuevoTutor, Siguiente: nil, Anterior: nil}
	if lista.Tamanio == 0{
		lista.Inicio = nuevoNodo
		lista.Inicio.Siguiente = nuevoNodo
		lista.Inicio.Anterior = nuevoNodo
		lista.Tamanio++
	}else{
		aux := lista.Inicio
		contador := 1
		for contador < lista.Tamanio{
			if lista.Inicio.Tutor.Carnet > carnet{
				nuevoNodo.Siguiente = lista.Inicio
				nuevoNodo.Anterior = lista.Inicio.Anterior
				lista.Inicio.Anterior.Siguiente = nuevoNodo
				lista.Inicio.Anterior = nuevoNodo
				lista.Tamanio++
				return
			}
			if aux.Tutor.Carnet < carnet{
				aux = aux.Siguiente
			}else{
				nuevoNodo.Anterior = aux.Anterior
				aux.Anterior.Siguiente = nuevoNodo
				nuevoNodo.Siguiente = aux
				aux.Anterior = nuevoNodo
				lista.Tamanio++
				return
			}
			contador++
		}
		if aux.Tutor.Carnet > carnet {
			nuevoNodo.Anterior = aux.Anterior
			aux.Anterior.Siguiente = nuevoNodo
			nuevoNodo.Siguiente = aux
			aux.Anterior = nuevoNodo
			lista.Tamanio++
			return
		}
		nuevoNodo.Anterior = aux
		nuevoNodo.Siguiente = lista.Inicio
		aux.Siguiente = nuevoNodo
		lista.Inicio.Anterior = nuevoNodo
		lista.Tamanio++
	}
}

func (lista *ListaCircular) MostrarTutores(){
	aux := lista.Inicio
	contador := 0
	for contador < lista.Tamanio{
		fmt.Println(aux.Tutor.Curso, "->", aux.Tutor.Nombre)
		aux = aux.Siguiente
		contador++
	}
}

func (lista *ListaCircular) BuscarTutor(curso string) *NodoListaCircular{
	aux := lista.Inicio
	contador := 0
	for contador < lista.Tamanio{
		if aux.Tutor.Curso == curso{
			return aux
		}
		aux = aux.Siguiente
		contador++
	}
	return nil
}

func (lista *ListaCircular) BuscarCurso(curso string) bool {
	if lista.Tamanio == 0 {
		return false
	} else {
		aux := lista.Inicio
		contador := 0
		for contador < lista.Tamanio {
			if aux.Tutor.Curso == curso {
				return true
			}
			aux = aux.Siguiente
			contador++
		}
	}
	return false
}

func (lista *ListaCircular) ReporteTutores() {
	nombre_archivo := "./listadoblecircular.dot"
	nombre_imagen := "./listadoblecircular.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	aux := lista.Inicio
	contador := 0
	for i := 0; i < lista.Tamanio; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + strconv.Itoa(aux.Tutor.Carnet) + "\"];\n"
		aux = aux.Siguiente
	}
	for i := 0; i < (lista.Tamanio - 1); i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodo0 \n"
	texto += "nodo0 -> " + "nodo" + strconv.Itoa(contador) + "\n"
	texto += "}"
	Archivos.CrearArchivo(nombre_archivo)
	Archivos.EscribirArchivo(texto, nombre_archivo)
	Archivos.Ejecutar(nombre_imagen, nombre_archivo)
}