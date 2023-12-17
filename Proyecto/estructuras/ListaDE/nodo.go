package ListaDE

type Estudiante struct {
	Carnet int
	Nombre string
}

type NodoListaDoble struct {
	Estudiante *Estudiante
	Siguiente  *NodoListaDoble
	Anterior   *NodoListaDoble
}