package ArbolAVL

import (
	"Proyecto/estructuras/Archivos"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Arbol struct {
	Raiz *NodoArbol
}

//Función para leer el archivo json con los datos de los cursos y agregarlos al árbol AVL
func (arbol *Arbol) LeerJsonCursos(ruta string) {
	_, err := os.Open(ruta)
	if err != nil{
		fmt.Println("[ERROR]: No se pudo abrir el archivo")
		return
	}
	contenido, err := os.ReadFile(ruta)
	if err != nil {
		log.Fatal("[ERROR] Problemas al leer el archivo:", err)
	}
	var datos DatosCursos
	//Se guardan los datos del archivo en un arreglo llamado datos
	err = json.Unmarshal(contenido, &datos)
	if err != nil {
		log.Fatal("[ERROR] Problemas al decodificar el JSON:", err)
	}
	//Se recorre el arreglo de cursos y se insertan al árbol AVL
	for _, curso := range datos.Cursos {
		arbol.InsertarCurso(curso.Codigo)
	}
	fmt.Println("Cursos cargados con éxito")
}

//Función que crea un nuevo nodo con el dato enviado como parámetro y lo inserta en el árbol
func (arbol *Arbol) InsertarCurso(dato string) {
	nuevoNodo := &NodoArbol{Dato: dato}
	arbol.Raiz = arbol.insertarNodo(arbol.Raiz, nuevoNodo)
}

/*
Función para insertar un nuevo nodo al árbol AVL tomando en cuenta el factor de desbalance 
del mismo, y si es necesario realizar rotaciones a la izquierda y/o derecha.
*/
func (arbol *Arbol) insertarNodo(raiz *NodoArbol, nuevoNodo *NodoArbol) *NodoArbol {
	if raiz == nil {
		raiz = nuevoNodo
	} else {
		/*
		IF -> Si el dato en la raiz es mayor al dato que se va a insertar, recorrer el subárbol izquierdo
		ELSE -> Si el dato en la raiz es menor al dato que se va a insertar, recorrer en el subárbol derecho
		*/
		if raiz.Dato > nuevoNodo.Dato {
			raiz.Izquierda = arbol.insertarNodo(raiz.Izquierda, nuevoNodo)
		} else {
			raiz.Derecha = arbol.insertarNodo(raiz.Derecha, nuevoNodo)
		}
	}
	//Variable con la altura más grande entre la altura del subárbol izquierdo y la del derecho
	numeroMax := math.Max(float64(arbol.altura(raiz.Izquierda)), float64(arbol.altura(raiz.Derecha)))
	raiz.Altura = 1 + int(numeroMax)
	balanceo := arbol.equilibrio(raiz)
	raiz.Factor_Equilibrio = balanceo
	if balanceo > 1 && nuevoNodo.Dato > raiz.Derecha.Dato {
		return arbol.rotacionIzquierda(raiz)
	} else if balanceo < -1 && nuevoNodo.Dato < raiz.Izquierda.Dato {
		return arbol.rotacionDerecha(raiz)
	} else if balanceo > 1 && nuevoNodo.Dato < raiz.Derecha.Dato {
		raiz.Derecha = arbol.rotacionDerecha(raiz.Derecha)
		return arbol.rotacionIzquierda(raiz)
	} else if balanceo < -1 && nuevoNodo.Dato > raiz.Izquierda.Dato {
		raiz.Izquierda = arbol.rotacionIzquierda(raiz.Izquierda)
		return arbol.rotacionDerecha(raiz)
	}
	return raiz
}

//Función para retornar la altura de un nodo
func (arbol *Arbol) altura(raiz *NodoArbol) int {
	if raiz == nil {
		return 0
	}
	return raiz.Altura
}

/*
Función para determinar el factor de equilibrio del nodo enviado como parámetro
FACTOR_EQUILIBRIO = Altura del subárbol derecho - Altura del subárbol izquierdo
*/
func (arbol *Arbol) equilibrio(raiz *NodoArbol) int {
	if raiz == nil {
		return 0
	}
	return (arbol.altura(raiz.Derecha) - arbol.altura(raiz.Izquierda))
}

/*
Función para realizar la rotación a la izquierda y de esa forma balancear el árbol AVL
	10						15
		15			->	10		20
			20
*/
func (arbol *Arbol) rotacionIzquierda(raiz *NodoArbol) *NodoArbol { //raiz = 10
	raiz_derecho := raiz.Derecha //raiz_derecho = 15
	hijo_izquierdo := raiz_derecho.Izquierda //hijo_izquierdo = nil (nodo que exista a la izquierda del 15)
	raiz_derecho.Izquierda = raiz //raiz_derecho.Izquierda = 10
	raiz.Derecha = hijo_izquierdo //raiz.Derecha = nil (nodo que estaba a la izquierda del 15)
	//Calcular nuevamente las alturas de raiz
	numeroMax := math.Max(float64(arbol.altura(raiz.Izquierda)), float64(arbol.altura(raiz.Derecha)))
	raiz.Altura = 1 + int(numeroMax)
	raiz.Factor_Equilibrio = arbol.equilibrio(raiz)
	//Calcular nuevamente las alturas de raiz.derecho
	numeroMax = math.Max(float64(arbol.altura(raiz_derecho.Izquierda)), float64(arbol.altura(raiz_derecho.Derecha)))
	raiz_derecho.Altura = 1 + int(numeroMax)
	raiz_derecho.Factor_Equilibrio = arbol.equilibrio(raiz_derecho)
	return raiz_derecho
}

/*
Función para realizar la rotación a la derecha y de esa forma balancear el árbol AVL
			20				15
		15		  ->	10		20
	10
*/
func (arbol *Arbol) rotacionDerecha(raiz *NodoArbol) *NodoArbol { //raiz = 20
	raiz_izquierdo := raiz.Izquierda //raiz_izquierdo = 15
	hijo_derecho := raiz_izquierdo.Derecha //hijo_derecho = nil (nodo que exista a la derecha del 15)
	raiz_izquierdo.Derecha = raiz //raiz_izquierdo.Derecha = 20
	raiz.Izquierda = hijo_derecho //raiz.Izquierda = nil (nodo que estaba a la derecha del 15)
	numeroMax := math.Max(float64(arbol.altura(raiz.Izquierda)), float64(arbol.altura(raiz.Derecha)))
	raiz.Altura = 1 + int(numeroMax)
	raiz.Factor_Equilibrio = arbol.equilibrio(raiz)
	numeroMax = math.Max(float64(arbol.altura(raiz_izquierdo.Izquierda)), float64(arbol.altura(raiz_izquierdo.Derecha)))
	raiz_izquierdo.Altura = 1 + int(numeroMax)
	raiz_izquierdo.Factor_Equilibrio = arbol.equilibrio(raiz_izquierdo)
	return raiz_izquierdo
}

/*
Función para comprobar si existe un curso determinado en el árbol AVL,
si este existe, retorna true, caso contrario retorna false
*/
func (arbol *Arbol) BuscarCurso(dato string) bool {
	buscar_curso := arbol.buscarNodo(dato, arbol.Raiz)
	if buscar_curso != nil {
		return true
	}
	return false
}

//Función para buscar un curso determinado dentro del árbol AVL y retornar el nodo que lo contiene
func (arbol *Arbol) buscarNodo(dato string, raiz *NodoArbol) *NodoArbol {
	var curso_encontrado *NodoArbol
	if raiz != nil {
		if raiz.Dato == dato {
			curso_encontrado = raiz
		} else {
			/*
			IF -> Si el dato en la raiz es mayor al dato que se busca, buscar en el subárbol izquierdo
			ELSE -> Si el dato en la raiz es menor al que se busca, buscar en el subárbol derecho
			*/
			if raiz.Dato > dato {
				curso_encontrado = arbol.buscarNodo(dato, raiz.Izquierda)
			} else {
				curso_encontrado = arbol.buscarNodo(dato, raiz.Derecha)
			}
		}
	}
	return curso_encontrado
}

//Función que genera el reporte de cursos de acuerdo al árbol AVL
func (arbol *Arbol) ReporteCursos() {
	cadena := ""
	nombre_archivo := "./ArbolAVL.dot"
	nombre_imagen := "./ArbolAVL.jpg"
	if arbol.Raiz != nil {
		cadena += "digraph arbol{ "
		cadena += arbol.retornarValoresArbol(arbol.Raiz, 0)
		cadena += "}"
	}
	Archivos.CrearArchivo(nombre_archivo)
	Archivos.EscribirArchivo(cadena, nombre_archivo)
	Archivos.Ejecutar(nombre_imagen, nombre_archivo)
}

//Función donde se establecen los nodos y sus conexiones para la generación del reporte
func (arbol *Arbol) retornarValoresArbol(raiz *NodoArbol, indice int) string {
	cadena := ""
	numero := indice + 1
	if raiz != nil {
		cadena += "\""
		cadena += raiz.Dato
		cadena += "\" ;"
		if raiz.Izquierda != nil && raiz.Derecha != nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Dato
			cadena += "\" -> "
			cadena += arbol.retornarValoresArbol(raiz.Izquierda, numero)
			cadena += "\""
			cadena += raiz.Dato
			cadena += "\" -> "
			cadena += arbol.retornarValoresArbol(raiz.Derecha, numero)
			cadena += "{rank=same" + "\"" + (raiz.Izquierda.Dato) + "\"" + " -> " + "\"" + (raiz.Derecha.Dato) + "\"" + " [style=invis]}; "
		} else if raiz.Izquierda != nil && raiz.Derecha == nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Dato
			cadena += "\" -> "
			cadena += arbol.retornarValoresArbol(raiz.Izquierda, numero)
			cadena += "\""
			cadena += raiz.Dato
			cadena += "\" -> "
			cadena += "x" + strconv.Itoa(numero) + "[style=invis]"
			cadena += "{rank=same" + "\"" + (raiz.Izquierda.Dato) + "\"" + " -> " + "x" + strconv.Itoa(numero) + " [style=invis]}; "
		} else if raiz.Izquierda == nil && raiz.Derecha != nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Dato
			cadena += "\" -> "
			cadena += "x" + strconv.Itoa(numero) + "[style=invis]"
			cadena += "; \""
			cadena += raiz.Dato
			cadena += "\" -> "
			cadena += arbol.retornarValoresArbol(raiz.Derecha, numero)
			cadena += "{rank=same" + " x" + strconv.Itoa(numero) + " -> \"" + (raiz.Derecha.Dato) + "\"" + " [style=invis]}; "
		}
	}
	return cadena
}