package ListaCDE

type Tutor struct {
	Carnet int
	Nombre string
	Curso  string
	Nota   int
}

type NodoListaCircular struct {
	Tutor     *Tutor
	Siguiente *NodoListaCircular
	Anterior  *NodoListaCircular
}