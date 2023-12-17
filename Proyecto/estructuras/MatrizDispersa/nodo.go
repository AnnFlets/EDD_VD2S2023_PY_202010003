package MatrizDispersa

type Dato struct {
	Carnet_Tutor      int
	Carnet_Estudiante int
	Curso             string
}

type NodoMatriz struct {
	Siguiente *NodoMatriz
	Anterior  *NodoMatriz
	Abajo     *NodoMatriz
	Arriba    *NodoMatriz
	PosX      int
	PosY      int
	Dato      *Dato
}