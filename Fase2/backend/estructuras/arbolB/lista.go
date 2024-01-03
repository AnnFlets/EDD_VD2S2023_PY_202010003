package arbolB

type NodoLista struct {
	Tutor     *NodoB
	Siguiente *NodoLista
}

type ListaSimple struct {
	Inicio   *NodoLista
	Longitud int
}

// Función para insertar un tutor enviado como parámetro en la lista enlazada simple
func (lista *ListaSimple) InsertarTutor(tutor *NodoB) {
	if lista.Longitud == 0 {
		nuevo := &NodoLista{Tutor: tutor, Siguiente: nil}
		lista.Inicio = nuevo
		lista.Longitud++
	} else {
		aux := lista.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = &NodoLista{Tutor: tutor, Siguiente: nil}
		lista.Longitud++
	}
}