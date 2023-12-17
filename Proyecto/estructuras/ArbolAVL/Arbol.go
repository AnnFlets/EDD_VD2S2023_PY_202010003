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
	err = json.Unmarshal(contenido, &datos)
	if err != nil {
		log.Fatal("[ERROR] Problemas al decodificar el JSON:", err)
	}
	for _, curso := range datos.Cursos {
		arbol.InsertarCurso(curso.Codigo)
	}
	fmt.Println("Cursos cargados con éxito")
}

func (arbol *Arbol) InsertarCurso(dato string) {
	nuevoNodo := &NodoArbol{Dato: dato}
	arbol.Raiz = arbol.insertarNodo(arbol.Raiz, nuevoNodo)
}

func (arbol *Arbol) insertarNodo(raiz *NodoArbol, nuevoNodo *NodoArbol) *NodoArbol {
	if raiz == nil {
		raiz = nuevoNodo
	} else {
		if raiz.Dato > nuevoNodo.Dato {
			raiz.Izquierda = arbol.insertarNodo(raiz.Izquierda, nuevoNodo)
		} else {
			raiz.Derecha = arbol.insertarNodo(raiz.Derecha, nuevoNodo)
		}
	}
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

func (arbol *Arbol) altura(raiz *NodoArbol) int {
	if raiz == nil {
		return 0
	}
	return raiz.Altura
}

func (arbol *Arbol) equilibrio(raiz *NodoArbol) int {
	if raiz == nil {
		return 0
	}
	return (arbol.altura(raiz.Derecha) - arbol.altura(raiz.Izquierda))
}

func (arbol *Arbol) rotacionIzquierda(raiz *NodoArbol) *NodoArbol {
	raiz_derecho := raiz.Derecha
	hijo_izquierdo := raiz_derecho.Izquierda
	raiz_derecho.Izquierda = raiz
	raiz.Derecha = hijo_izquierdo
	numeroMax := math.Max(float64(arbol.altura(raiz.Izquierda)), float64(arbol.altura(raiz.Derecha)))
	raiz.Altura = 1 + int(numeroMax)
	raiz.Factor_Equilibrio = arbol.equilibrio(raiz)
	numeroMax = math.Max(float64(arbol.altura(raiz_derecho.Izquierda)), float64(arbol.altura(raiz_derecho.Derecha)))
	raiz_derecho.Altura = 1 + int(numeroMax)
	raiz_derecho.Factor_Equilibrio = arbol.equilibrio(raiz_derecho)
	return raiz_derecho
}

func (arbol *Arbol) rotacionDerecha(raiz *NodoArbol) *NodoArbol {
	raiz_izquierdo := raiz.Izquierda
	hijo_derecho := raiz_izquierdo.Derecha
	raiz_izquierdo.Derecha = raiz
	raiz.Izquierda = hijo_derecho
	numeroMax := math.Max(float64(arbol.altura(raiz.Izquierda)), float64(arbol.altura(raiz.Derecha)))
	raiz.Altura = 1 + int(numeroMax)
	raiz.Factor_Equilibrio = arbol.equilibrio(raiz)
	numeroMax = math.Max(float64(arbol.altura(raiz_izquierdo.Izquierda)), float64(arbol.altura(raiz_izquierdo.Derecha)))
	raiz_izquierdo.Altura = 1 + int(numeroMax)
	raiz_izquierdo.Factor_Equilibrio = arbol.equilibrio(raiz_izquierdo)
	return raiz_izquierdo
}

func (arbol *Arbol) BuscarCurso(dato string) bool {
	buscar_curso := arbol.buscarNodo(dato, arbol.Raiz)
	if buscar_curso != nil {
		return true
	}
	return false
}

func (arbol *Arbol) buscarNodo(dato string, raiz *NodoArbol) *NodoArbol {
	var curso_encontrado *NodoArbol
	if raiz != nil {
		if raiz.Dato == dato {
			curso_encontrado = raiz
		} else {
			if raiz.Dato > dato {
				curso_encontrado = arbol.buscarNodo(dato, raiz.Izquierda)
			} else {
				curso_encontrado = arbol.buscarNodo(dato, raiz.Derecha)
			}
		}
	}
	return curso_encontrado
}

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