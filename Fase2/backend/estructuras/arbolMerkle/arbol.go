package arbolMerkle

import (
	"Fase2/estructuras/generarArchivos"
	"encoding/hex"
	"math"
	"strconv"
	"time"

	"golang.org/x/crypto/sha3"
)

type ArbolMerkle struct {
	Raiz *NodoMerkle
	BloqueDatos *NodoBloqueDatos
	CantidadBloques int
}

//Función para obtener la fecha actual en el formato DD-MM-YYYY::HH:MM:SS
func fechaHoraActual() string {
	fecha_hora := time.Now()
	formato := "02-01-2006::15:04:05"
	fecha_hora_formato := fecha_hora.Format(formato)
	return fecha_hora_formato
}

//Función para crear el árbol de Merkle con los datos almacenados en los data blocks
func (arbol *ArbolMerkle) GenerarArbol() {
	nivel := 1
	//Se ejecuta el ciclo hasta que la cantidad de bloques del árbol sea menor a los valores de la potencia de 2 que se van obteniendo
	for int(math.Pow(2, float64(nivel))) < arbol.CantidadBloques {
		nivel++
	}
	//Se ejecuta el ciclo la cantidad de veces asociada a la cantidad de bloques de datos faltantes para que sea un resultado de una potencia de 2
	for i := arbol.CantidadBloques; i < int(math.Pow(2, float64(nivel))); i++ {
		arbol.AgregarBloque(strconv.Itoa(i+1), "nulo", 0)
	}
	arbol.generarHash()
}

//Función para agregar un bloque de datos a la lista doblemente enlazada del árbol de Merkle.
func (arbol *ArbolMerkle) AgregarBloque(accion string, nombre string, carnet_tutor int) {
	nuevo_registro := &InformacionBloque{FechaHora: fechaHoraActual(), Accion: accion, Nombre: nombre, Tutor: carnet_tutor}
	nuevo_bloque := &NodoBloqueDatos{Valor: nuevo_registro}
	/*
	IF -> No hay elementos en la lista
	ELSE -> Hay elementos en la lista
	*/
	if arbol.BloqueDatos == nil {
		arbol.BloqueDatos = nuevo_bloque
		arbol.CantidadBloques++
	} else {
		aux := arbol.BloqueDatos
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		nuevo_bloque.Anterior = aux
		aux.Siguiente = nuevo_bloque
		arbol.CantidadBloques++
	}
}

//Función para encriptar la información de los data blocks con SHA3, relacionar los nodos hojas del árbol Merkle con los bloques de datos y designarles el hash correspondiente
func (arbol *ArbolMerkle) generarHash() {
	var arrayNodos []*NodoMerkle
	aux := arbol.BloqueDatos
	//Se recorren los data block, se concatena la información de estos, se encripta con SHA3, se crea el nodo_hoja que contendrá el hash del bloque de datos correspondiente y al arreglo de nodos se insertan los nodos hojas que se creen. 
	for aux != nil {
		concatenacion := aux.Valor.FechaHora + aux.Valor.Accion + aux.Valor.Nombre + strconv.Itoa(aux.Valor.Tutor)
		encriptado := arbol.encriptarSHA3(concatenacion)
		nodo_hoja := &NodoMerkle{Valor: encriptado, Bloque: aux}
		arrayNodos = append(arrayNodos, nodo_hoja)
		aux = aux.Siguiente
	}
	arbol.Raiz = arbol.crearArbol(arrayNodos)
}

//Función para encriptar con SHA3 la cadena recibida como parámetro y retornar el resultado de dicho proceso
func (arbol *ArbolMerkle) encriptarSHA3(cadena string) string {
	hash := sha3.New256()
	hash.Write([]byte(cadena))
	encriptacion := hex.EncodeToString(hash.Sum(nil))
	return encriptacion
}

//Función para formar el árbol, generando el hash de la concatenación de los valores de 2 nodos hermanos y asignándosela a un nodo raíz. Retorna la raíz del árbol
func (arbol *ArbolMerkle) crearArbol(arrayNodos []*NodoMerkle) *NodoMerkle {
	var auxNodos []*NodoMerkle
	var raiz *NodoMerkle
	/*
	IF -> Si solo hay 2 nodos en el arreglo (que serían los nodos hijos de la raíz)
	ELSE -> Si hay más nodos, y por tanto no son los hijos directos de la raíz
	*/
	if len(arrayNodos) == 2 {
		encriptado := arbol.encriptarSHA3(arrayNodos[0].Valor + arrayNodos[1].Valor)
		raiz = &NodoMerkle{Valor: encriptado}
		raiz.Izquierda = arrayNodos[0]
		raiz.Derecha = arrayNodos[1]
		return raiz
	} else {
		for i := 0; i < len(arrayNodos); i += 2 {
			encriptado := arbol.encriptarSHA3(arrayNodos[i].Valor + arrayNodos[i+1].Valor)
			nodoRaiz := &NodoMerkle{Valor: encriptado}
			nodoRaiz.Izquierda = arrayNodos[i]
			nodoRaiz.Derecha = arrayNodos[i+1]
			auxNodos = append(auxNodos, nodoRaiz)
		}
		return arbol.crearArbol(auxNodos)
	}
}

//Función para generar el reporte del árbol de Merkle en .dot y .jpg
func (arbol *ArbolMerkle) ReporteMerkleLibros(nombre string) {
	cadena := ""
	nombre_archivo := "./Reporte/arbolMerkle.dot"
	nombre_imagen := "./Reporte/arbolMerkle.jpg"
	if arbol.Raiz != nil {
		cadena += "digraph arbol { node [shape=box];"
		cadena += arbol.retornarValoresArbol(arbol.Raiz, 0)
		cadena += "}"
	}
	generarArchivos.CrearArchivo(nombre_archivo)
	generarArchivos.EscribirArchivo(cadena, nombre_archivo)
	generarArchivos.Ejecutar(nombre_imagen, nombre_archivo)
}

func (arbol *ArbolMerkle) retornarValoresArbol(raiz *NodoMerkle, indice int) string {
	cadena := ""
	numero := indice + 1
	if raiz != nil {
		cadena += "\""
		cadena += raiz.Valor[:20]
		cadena += "\" [dir=back];\n"
		if raiz.Izquierda != nil && raiz.Derecha != nil {
			cadena += "\""
			cadena += raiz.Valor[:20]
			cadena += "\" -> "
			cadena += arbol.retornarValoresArbol(raiz.Izquierda, numero)
			cadena += "\""
			cadena += raiz.Valor[:20]
			cadena += "\" -> "
			cadena += arbol.retornarValoresArbol(raiz.Derecha, numero)
			cadena += "{rank=same" + "\"" + (raiz.Izquierda.Valor[:20]) + "\"" + " -> " + "\"" + (raiz.Derecha.Valor[:20]) + "\"" + " [style=invis]}; \n"
		}
	}
	if raiz.Bloque != nil {
		cadena += "\""
		cadena += raiz.Valor[:20]
		cadena += "\" -> "
		cadena += "\""
		cadena += raiz.Bloque.Valor.FechaHora + "\n" + raiz.Bloque.Valor.Accion + "\n" + raiz.Bloque.Valor.Nombre + "\n" + strconv.Itoa(raiz.Bloque.Valor.Tutor)
		cadena += "\" [dir=back];\n "
	}
	return cadena
}