package arbolB

type Libro struct {
	Nombre    string
	Contenido string
	Estado    int
	Curso     string
	Tutor     int
}
type Publicacion struct {
	Contenido string
	Curso     string
}

type Tutor struct {
	Carnet        int
	Nombre        string
	Curso         string
	Password      string
	Libros        []*Libro
	Publicaciones []*Publicacion
}

type NodoB struct {
	Valor *Tutor
	//Punteros para movernos en el arreglo (En el nodo, de izquierda a derecha)
	Siguiente *NodoB
	Anterior  *NodoB
	//Punteros para definir los hijos izquierdo y derecho del elemento del nodo
	Izquierdo *RamaB
	Derecho   *RamaB
}