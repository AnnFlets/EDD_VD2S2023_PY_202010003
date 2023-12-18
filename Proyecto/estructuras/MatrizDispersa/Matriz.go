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

//Función que inserta el elemento correspondiente a la asignación en la matriz dispersa
func (matriz *Matriz) InsertarElemento(carnet_estudiante int, carnet_tutor int, curso string) {
	nodo_columna := matriz.buscarColumna(carnet_tutor, curso)
	nodo_fila := matriz.buscarFila(carnet_estudiante)
	/*
	IF -> Si la columna y la fila no se encontraron (no existen)
	ELSE IF 1 -> Si la columna existe y la fila no existe
	ELSE IF 2 -> Si la columna no existe y la fila sí existe
	ELSE IF 3 -> Si la columna y la fila sí existen
	ELSE -> Otro caso
	*/
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

/*
Función que recorre el encabezado de las columnas, empezando desde la raiz hacia la derecha 
y busca si algún nodo almacena el mismo carnet de un tutor y curso a los enviados como parámetros.
Retorna el nodo con datos iguales si este es encontrado, de lo contrario retorna nil
*/
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

/*
Función que recorre el encabezado de las filas, empezando desde la raiz hacia abajo y busca
si algún nodo almacena el mismo carnet de estudiante al enviado como parámetro.
Retorna el nodo con dato igual si este es encontrado, de lo contrario retorna nil
*/
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

/*
Función que crea el nodo en el encabezado de las columnas, recibiendo como uno de los parámetros
la posición en "x" que debe ocupar el nodo (la cual corresponde a la cantidad de tutores
existentes en el encabezado), su posición en "y" siempre será -1 por ser el encabezado.
Luego inserta el nodo en la columna de la matriz y retorna el nodo.
*/
func (matriz *Matriz) nuevaColumna(x int, carnet_tutor int, curso string) *NodoMatriz {
	nuevoNodo := &NodoMatriz{PosX: x, PosY: -1, Dato: &Dato{Carnet_Tutor: carnet_tutor, Carnet_Estudiante: 0, Curso: curso}}
	columna := matriz.insertarColumna(nuevoNodo, matriz.Raiz)
	return columna
}

/*
Función que crea el nodo en el encabezado de las filas, recibiendo como uno de los parámetros
la posición en "y" que debe ocupar el nodo (la cual corresponde a la cantidad de estudiantes
existentes en el encabezado), su posición en "x" siempre será -1 por ser el encabezado.
Luego inserta el nodo en la fila de la matriz y retorna el nodo.
*/
func (matriz *Matriz) nuevaFila(y int, carnet_estudiante int, curso string) *NodoMatriz {
	nuevoNodo := &NodoMatriz{PosX: -1, PosY: y, Dato: &Dato{Carnet_Tutor: 0, Carnet_Estudiante: carnet_estudiante, Curso: curso}}
	fila := matriz.insertarFila(nuevoNodo, matriz.Raiz)
	return fila
}

/*
Función para insertar un nodo en la matriz dispersa, asignando los apuntadores para las columnas
"Siguiente" y "Anterior" ("Derecha" e "Izquierda")
*/
func (matriz *Matriz) insertarColumna(nuevoNodo *NodoMatriz, nodoRaiz *NodoMatriz) *NodoMatriz {
	//Variable temporal que contiene el nodo encabezado de la fila o la raiz de la matriz
	temp := nodoRaiz
	//Para identificar si se va a insertar ordenado o no
	piv := false
	//Se recorren las columnas de temp hacia la derecha
	for {
		/*
		IF -> Si la posición en "x" de temp es igual a la posición en "x" del nodo a insertar
		ELSE IF -> Si la posición en "x" de temp es mayor a la posición en "x" del nodo a insertar
		*/
		if temp.PosX == nuevoNodo.PosX {
			/*
			Se sustituyen la posición en "y" y el dato guardados en el temporal por los del nodo a insertar
			*/
			temp.PosY = nuevoNodo.PosY
			temp.Dato = nuevoNodo.Dato
			return temp
		} else if temp.PosX > nuevoNodo.PosX {
			piv = true
			break
		}
		/*
		Esta sería la condición del for para recorrer las columnas
		IF -> Si el nodo a la derecha de temp es distinto de nil, seguir recorriendo
		ELSE -> Si el nodo a la derecha es nil, parar el ciclo
		*/
		if temp.Siguiente != nil {
			temp = temp.Siguiente
		} else {
			break
		}
	}
	/*
	IF -> para insertar un nodo de forma ordenada asignando sus apuntadores de la Izquierda y Derecha (insertar entre nodos)
	ELSE -> para insertar el nodo (al final)
	*/
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

/*
Función para insertar un nodo en la matriz dispersa, asignando los apuntadores para las filas
"Arriba" y "Abajo"
*/
func (matriz *Matriz) insertarFila(nuevoNodo *NodoMatriz, nodoRaiz *NodoMatriz) *NodoMatriz {
	//Variable temporal que contiene el nodo encabezado de la columna o la raiz de la matriz
	temp := nodoRaiz
	//Para identificar si se va a insertar ordenado o no
	piv := false
	//Se recorren las filas de temp hacia abajo
	for {
		/*
		IF -> Si la posición en "y" de temp es igual a la posición en "y" del nodo a insertar
		ELSE IF -> Si la posición en "y" de temp es mayor a la posición en "y" del nodo a insertar
		*/
		if temp.PosY == nuevoNodo.PosY {
			/*
			Se sustituyen la posición en "x" y el dato guardados en el temporal por los del nodo a insertar
			*/
			temp.PosX = nuevoNodo.PosX
			temp.Dato = nuevoNodo.Dato
			return temp
		} else if temp.PosY > nuevoNodo.PosY {
			piv = true
			break
		}
		/*
		Esta sería la condición del for para recorrer las filas
		IF -> Si el nodo abajo de temp es distinto de nil, seguir recorriendo
		ELSE -> Si el nodo de abajo es nil, parar el ciclo
		*/
		if temp.Abajo != nil {
			temp = temp.Abajo
		} else {
			break
		}
	}
	/*
	IF -> para insertar un nodo de forma ordenada asignando sus apuntadores de la Arriba y Abajo (insertar entre nodos)
	ELSE -> para insertar el nodo (al final)
	*/
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

//Función que genera el reporte de asignaciones de acuerdo a la matriz dispersa
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