package ArbolAVL

type Curso struct {
	Codigo string `json:"Codigo"`
	Nombre string `json:"Nombre"`
}

type DatosCursos struct {
	Cursos []Curso `json:"Cursos"`
}

type NodoArbol struct {
	Izquierda         *NodoArbol
	Derecha           *NodoArbol
	Dato              string
	Altura            int
	Factor_Equilibrio int
}