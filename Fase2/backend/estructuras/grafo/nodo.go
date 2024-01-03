package grafo

type NodoGrafo struct {
	Siguiente *NodoGrafo
	Abajo     *NodoGrafo
	Curso     string
}

type PeticionGrafo struct {
	NombreArchivo string
}