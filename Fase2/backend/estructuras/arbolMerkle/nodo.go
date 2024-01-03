package arbolMerkle

//Información de los Data Blocks
type InformacionBloque struct {
	FechaHora string
	Accion    string
	Nombre    string
	Tutor     int
}

//Estructura para los nodos hoja (Data Blocks, en donde se tendrá una lista enlazada doble)
type NodoBloqueDatos struct {
	Siguiente *NodoBloqueDatos
	Anterior  *NodoBloqueDatos
	Informacion *InformacionBloque
}

//Estructura para los nodos del árbol de Merkle
type NodoMerkle struct {
	Izquierda *NodoMerkle
	Derecha   *NodoMerkle
	Bloque    *NodoBloqueDatos
	Valor     string
}