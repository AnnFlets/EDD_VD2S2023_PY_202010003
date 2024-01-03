package arbolB

type RamaB struct {
	//Primer elemento que existe en cada rama
	Primero *NodoB
	//Para identificar si la rama es hoja o no
	Hoja bool
	//Saber cuántos elementos hay en la rama y si es necesario hacer una división o no
	Contador int
}

// Función para insertar de forma ordenada un nuevo nodo en la rama
func (rama *RamaB) InsertarNodo(nuevoNodo *NodoB) {
	/*
		IF -> Si el primer nodo de la rama está vacío, definir que el Primero de la rama es el nuevoNodo
		ELSE -> Si el primer nodo de la rama no es nulo
	*/
	if rama.Primero == nil {
		rama.Primero = nuevoNodo
	} else {
		/*
			IF -> Si el curso del nuevoNodo a insertar es menor que el del primer nodo de la rama
			ELSE IF 1 -> Si el segundo nodo de la rama es distinto de nulo (está ocupado)
			ELSE IF 2 -> Si el segundo nodo de la rama es igual a nulo
		*/
		if nuevoNodo.Tutor.Curso < rama.Primero.Tutor.Curso {
			nuevoNodo.Siguiente = rama.Primero
			rama.Primero.Izquierdo = nuevoNodo.Derecho
			rama.Primero.Anterior = nuevoNodo
			rama.Primero = nuevoNodo
		} else if rama.Primero.Siguiente != nil {
			/*
				IF -> Si el curso del segundo nodo de la rama es mayor al curso del nuevoNodo a insertar (insertar entre nodos)
				ELSE -> Si el curso del segundo nodo de la rama es menor o igual al curso del nuevoNodo a insertar (insertar al final)
			*/
			if rama.Primero.Siguiente.Tutor.Curso > nuevoNodo.Tutor.Curso {
				nuevoNodo.Siguiente = rama.Primero.Siguiente
				nuevoNodo.Anterior = rama.Primero
				rama.Primero.Siguiente.Izquierdo = nuevoNodo.Derecho
				rama.Primero.Derecho = nuevoNodo.Izquierdo
				rama.Primero.Siguiente.Anterior = nuevoNodo
				rama.Primero.Siguiente = nuevoNodo
			} else {
				aux := rama.Primero.Siguiente
				nuevoNodo.Anterior = aux
				aux.Derecho = nuevoNodo.Izquierdo
				aux.Siguiente = nuevoNodo
			}
		} else if rama.Primero.Siguiente == nil {
			nuevoNodo.Anterior = rama.Primero
			rama.Primero.Derecho = nuevoNodo.Izquierdo
			rama.Primero.Siguiente = nuevoNodo
		}
	}
	rama.Contador++
}