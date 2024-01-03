package grafo

type NodoGrafo struct {
	Siguiente *NodoGrafo
	Abajo     *NodoGrafo
	Valor     string
}

type PeticionGrafo struct {
	NombreArchivo string
}