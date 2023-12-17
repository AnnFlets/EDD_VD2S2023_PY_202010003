package MatrizDispersa

import (
	"Proyecto/estructuras/Archivos"
	"fmt"
	"strconv"
)

type Matriz struct {
	Raiz             *NodoMatriz
	Cantidad_Estudiantes int
	Cantidad_Tutores int
}

func (matriz *Matriz) InsertarElemento(carnet_estudiante int, carnet_tutor int, curso string) {
	nodo_columna := matriz.buscarColumna(carnet_tutor, curso)
	nodo_fila := matriz.buscarFila(carnet_estudiante)
	if nodo_columna == nil && nodo_fila == nil {
		nodo_columna = matriz.nuevaColumna(matriz.Cantidad_Tutores, carnet_tutor, curso)
		nodo_fila = matriz.nuevaFila(matriz.Cantidad_Estudiantes, carnet_estudiante, curso)
		matriz.Cantidad_Estudiantes++
		matriz.Cantidad_Tutores++
		nuevoNodo := &NodoMatriz{PosX: nodo_columna.PosX, PosY: nodo_fila.PosY, Dato: &Dato{Carnet_Tutor: carnet_tutor, Carnet_Estudiante: carnet_estudiante, Curso: curso}}
		nuevoNodo = matriz.insertarColumna(nuevoNodo, nodo_fila)
		nuevoNodo = matriz.insertarFila(nuevoNodo, nodo_columna)
	} else if nodo_columna != nil && nodo_fila == nil {
		nodo_fila = matriz.nuevaFila(matriz.Cantidad_Estudiantes, carnet_estudiante, curso)
		matriz.Cantidad_Estudiantes++
		nuevoNodo := &NodoMatriz{PosX: nodo_columna.PosX, PosY: nodo_fila.PosY, Dato: &Dato{Carnet_Tutor: carnet_tutor, Carnet_Estudiante: carnet_estudiante, Curso: curso}}
		nuevoNodo = matriz.insertarColumna(nuevoNodo, nodo_fila)
		nuevoNodo = matriz.insertarFila(nuevoNodo, nodo_columna)
	} else if nodo_columna == nil && nodo_fila != nil {
		nodo_columna = matriz.nuevaColumna(matriz.Cantidad_Tutores, carnet_tutor, curso)
		matriz.Cantidad_Tutores++
		nuevoNodo := &NodoMatriz{PosX: nodo_columna.PosX, PosY: nodo_fila.PosY, Dato: &Dato{Carnet_Tutor: carnet_tutor, Carnet_Estudiante: carnet_estudiante, Curso: curso}}
		nuevoNodo = matriz.insertarColumna(nuevoNodo, nodo_fila)
		nuevoNodo = matriz.insertarFila(nuevoNodo, nodo_columna)
	} else if nodo_columna != nil && nodo_fila != nil {
		nuevoNodo := &NodoMatriz{PosX: nodo_columna.PosX, PosY: nodo_fila.PosY, Dato: &Dato{Carnet_Tutor: carnet_tutor, Carnet_Estudiante: carnet_estudiante, Curso: curso}}
		nuevoNodo = matriz.insertarColumna(nuevoNodo, nodo_fila)
		nuevoNodo = matriz.insertarFila(nuevoNodo, nodo_columna)
	} else {
		fmt.Println("[ERROR]: No se pudo realizar la asignación")
	}
}

func (matriz *Matriz) buscarColumna(carnet_tutor int, curso string) *NodoMatriz {
	aux := matriz.Raiz
	for aux != nil {
		if aux.Dato.Carnet_Tutor == carnet_tutor && aux.Dato.Curso == curso {
			return aux
		}
		aux = aux.Siguiente
	}
	return nil
}

func (matriz *Matriz) buscarFila(carnet_estudiante int) *NodoMatriz {
	aux := matriz.Raiz
	for aux != nil {
		if aux.Dato.Carnet_Estudiante == carnet_estudiante {
			return aux
		}
		aux = aux.Abajo
	}
	return nil
}

func (matriz *Matriz) insertarColumna(nuevoNodo *NodoMatriz, nodoRaiz *NodoMatriz) *NodoMatriz {
	temp := nodoRaiz
	piv := false
	for {
		if temp.PosX == nuevoNodo.PosX {
			temp.PosY = nuevoNodo.PosY
			temp.Dato = nuevoNodo.Dato
			return temp
		} else if temp.PosX > nuevoNodo.PosX {
			piv = true
			break
		}
		if temp.Siguiente != nil {
			temp = temp.Siguiente
		} else {
			break
		}
	}
	if piv {
		nuevoNodo.Siguiente = temp
		nuevoNodo.Anterior = temp.Anterior
		temp.Anterior.Siguiente = nuevoNodo
		temp.Anterior = nuevoNodo
	} else {
		nuevoNodo.Anterior = temp
		temp.Siguiente = nuevoNodo
	}
	return nuevoNodo
}

func (matriz *Matriz) nuevaColumna(x int, carnet_Tutor int, curso string) *NodoMatriz {
	nuevoNodo := &NodoMatriz{PosX: x, PosY: -1, Dato: &Dato{Carnet_Tutor: carnet_Tutor, Carnet_Estudiante: 0, Curso: curso}}
	columna := matriz.insertarColumna(nuevoNodo, matriz.Raiz)
	return columna
}

func (matriz *Matriz) insertarFila(nuevoNodo *NodoMatriz, nodoRaiz *NodoMatriz) *NodoMatriz {
	temp := nodoRaiz
	piv := false
	for {
		if temp.PosY == nuevoNodo.PosY {
			temp.PosX = nuevoNodo.PosX
			temp.Dato = nuevoNodo.Dato
			return temp
		} else if temp.PosY > nuevoNodo.PosY {
			piv = true
			break
		}
		if temp.Abajo != nil {
			temp = temp.Abajo
		} else {
			break
		}
	}
	if piv {
		nuevoNodo.Abajo = temp
		nuevoNodo.Arriba = temp.Arriba
		temp.Arriba.Abajo = nuevoNodo
		temp.Arriba = nuevoNodo
	} else {
		nuevoNodo.Arriba = temp
		temp.Abajo = nuevoNodo
	}
	return nuevoNodo
}

func (matriz *Matriz) nuevaFila(y int, carnet_estudiante int, curso string) *NodoMatriz {
	nuevoNodo := &NodoMatriz{PosX: -1, PosY: y, Dato: &Dato{Carnet_Tutor: 0, Carnet_Estudiante: carnet_estudiante, Curso: curso}}
	fila := matriz.insertarFila(nuevoNodo, matriz.Raiz)
	return fila
}

func (matriz *Matriz) ReporteAsignaciones() {
	texto := ""
	nombre_archivo := "./matriz.dot"
	nombre_imagen := "./matriz.jpg"
	aux1 := matriz.Raiz
	aux2 := matriz.Raiz
	aux3 := matriz.Raiz
	if aux1 != nil {
		texto = "digraph MatrizCapa{ \n node[shape=box] \n rankdir=UD; \n {rank=min; \n"
		for aux1 != nil {
			texto += "nodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + "[label=\"" + strconv.Itoa(aux1.Dato.Carnet_Tutor) + "\" ,rankdir=LR,group=" + strconv.Itoa(aux1.PosX+1) + "]; \n"
			aux1 = aux1.Siguiente
		}
		texto += "}"
		aux2 = aux2.Abajo
		for aux2 != nil {
			aux1 = aux2
			texto += "{rank=same; \n"
			flag_raiz := true
			for aux1 != nil {
				if flag_raiz {
					texto += "nodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + "[label=\"" + strconv.Itoa(aux1.Dato.Carnet_Estudiante) + "\" ,group=" + strconv.Itoa(aux1.PosX+1) + "]; \n"
					flag_raiz = false
				} else {
					texto += "nodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + "[label=\"" + aux1.Dato.Curso + "\" ,group=" + strconv.Itoa(aux1.PosX+1) + "]; \n"
				}
				aux1 = aux1.Siguiente
			}
			texto += "}"
			aux2 = aux2.Abajo
		}
		aux2 = aux3
		for aux2 != nil {
			aux1 = aux2
			for aux1.Siguiente != nil {
				texto += "nodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + " -> " + "nodo" + strconv.Itoa(aux1.Siguiente.PosX+1) + strconv.Itoa(aux1.Siguiente.PosY+1) + " [dir=both];\n"
				aux1 = aux1.Siguiente
			}
			aux2 = aux2.Abajo
		}
		aux2 = aux3
		for aux2 != nil {
			aux1 = aux2
			for aux1.Abajo != nil {
				texto += "nodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + " -> " + "nodo" + strconv.Itoa(aux1.Abajo.PosX+1) + strconv.Itoa(aux1.Abajo.PosY+1) + " [dir=both];\n"
				aux1 = aux1.Abajo
			}
			aux2 = aux2.Siguiente
		}
		texto += "}"
	} else {
		texto = "No hay elementos en la matriz"
	}
	Archivos.CrearArchivo(nombre_archivo)
	Archivos.EscribirArchivo(texto, nombre_archivo)
	Archivos.Ejecutar(nombre_imagen, nombre_archivo)
}