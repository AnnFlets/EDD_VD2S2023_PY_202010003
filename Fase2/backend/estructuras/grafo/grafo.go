package grafo

import "Fase2/estructuras/generarArchivos"

type Grafo struct {
	Principal *NodoGrafo
}

// Función para insertar el valor de los cursos en una lista de adyacencia (útil para elaborar el grafo dirigido)
func (grafo *Grafo) InsertarCurso(curso string, post string) {
	/*
		IF -> Si el grafo está vacío
		ELSE -> Si hay elemento(s) en el grafo
	*/
	if grafo.Principal == nil {
		//Insertar el nuevo curso en una fila nueva
		grafo.insertarFila(curso)
		//Insertar los cursos post requisitos en la fila del curso determinado
		grafo.insertarColumna(curso, post)
	} else {
		grafo.insertarColumna(curso, post)
	}
}

// Función que inserta un curso, en una nueva fila, en la columna principal de la lista de adyacencia.
func (grafo *Grafo) insertarFila(curso string) {
	nuevoNodo := &NodoGrafo{Valor: curso}
	/*
		IF -> Si la lista de adyacencia está vacía
		ELSE -> Si la lista de adyacencia tiene elementos
	*/
	if grafo.Principal == nil {
		grafo.Principal = nuevoNodo
	} else {
		//Recorre la columna principal de la lista de adyacencia
		aux := grafo.Principal
		for aux.Abajo != nil {
			if aux.Valor == curso {
				return
			}
			aux = aux.Abajo
		}
		aux.Abajo = nuevoNodo
	}
}

// Función que inserta un curso post requisito en la lista enlazada de un curso determinado.
func (grafo *Grafo) insertarColumna(curso string, post string) {
	nuevoNodo := &NodoGrafo{Valor: post}
	/*
		IF -> Si la lista de adyacencia no está vacía y el curso del primer elemento es igual al curso recibido
		ELSE -> Si la lista de adyacencia está vacía o el curso del primer elemento no es igual al curso recibido
	*/
	if grafo.Principal != nil && grafo.Principal.Valor == curso {
		//Se inserta el curso post requisito en la columna principal de la lista de adyacencia
		grafo.insertarFila(post)
		aux := grafo.Principal
		//Se recorre la lista enlazada simple del elemento de la fila
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		//Se inserta en la lista el nodo con el curso post requisito
		aux.Siguiente = nuevoNodo
	} else {
		grafo.insertarFila(curso)
		aux := grafo.Principal
		//Se busca el curso en la columna principal de la lista de adyacencia
		for aux != nil {
			if aux.Valor == curso {
				break
			}
			aux = aux.Abajo
		}
		//Si se encontró el curso en la lista de adyacencia, se recorre la lista enlazada de dicho nodo y se inserta en la última posición
		if aux != nil {
			for aux.Siguiente != nil {
				aux = aux.Siguiente
			}
			aux.Siguiente = nuevoNodo
		}
	}
}

// Función para generar el reporte del grafo dirigido con los cursos guardados en .dot y .jpg
func (grafo *Grafo) ReporteGrafoCursos(nombre string) {
	cadena := ""
	nombre_archivo := "./Reporte/" + nombre + ".dot"
	nombre_imagen := "./Reporte/" + nombre + ".jpg"
	if grafo.Principal != nil {
		cadena += "digraph grafoDirigido{ \n rankdir=LR; \n node [shape=box]; layout=neato; \n nodo" + grafo.Principal.Valor + "[label=\"" + grafo.Principal.Valor + "\"]; \n"
		cadena += "node [shape = ellipse]; \n"
		cadena += grafo.retornarValoresMatriz()
		cadena += "\n}"
	}
	generarArchivos.CrearArchivo(nombre_archivo)
	generarArchivos.EscribirArchivo(cadena, nombre_archivo)
	generarArchivos.Ejecutar(nombre_imagen, nombre_archivo)
}

func (grafo *Grafo) retornarValoresMatriz() string {
	cadena := ""
	/*CREACION DE NODOS*/
	aux := grafo.Principal.Abajo //Filas
	aux1 := aux                  //Columnas
	/*CREACION DE NODOS CON LABELS*/
	for aux != nil {
		for aux1 != nil {
			cadena += "nodo" + aux1.Valor + "[label=\"" + aux1.Valor + "\" ]; \n"
			aux1 = aux1.Siguiente
		}
		if aux != nil {
			aux = aux.Abajo
			aux1 = aux
		}
	}
	/*CONEXION DE NODOS*/
	aux = grafo.Principal  //Filas
	aux1 = aux.Siguiente   //Columnas
	/*CREACION DE NODOS CON LABELS*/
	for aux != nil {
		for aux1 != nil {
			cadena += "nodo" + aux.Valor + " -> "
			cadena += "nodo" + aux1.Valor + "[len=1.00]; \n"
			aux1 = aux1.Siguiente
		}
		if aux.Abajo != nil {
			aux = aux.Abajo
			aux1 = aux.Siguiente
		} else {
			aux = aux.Abajo
		}
	}
	return cadena
}