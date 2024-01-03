package arbolB

import (
	"Fase2/estructuras/generarArchivos"
	"fmt"
	"strconv"
)

type ArbolB struct {
	Raiz  *RamaB
	Orden int
}
//Función para insertar un tutor en una rama del árbol B
func (arbol *ArbolB) InsertarTutor(carnet int, nombre string, curso string, password string) {
	tutor := &Tutor{Carnet: carnet, Nombre: nombre, Curso: curso, Password: password}
	nuevoNodo := &NodoB{Tutor: tutor}
	/*
	IF -> Si la raíz del árbol es nula, crear la rama raiz e insertar el nodo con los datos recibidos
	ELSE -> Si la raíz del árbol no es nula
	*/
	if arbol.Raiz == nil {
		arbol.Raiz = &RamaB{Primero: nil, Hoja: true, Contador: 0}
		arbol.Raiz.InsertarNodo(nuevoNodo)
	} else {
		obj := arbol.insertarRama(nuevoNodo, arbol.Raiz)
		//Si se hace una división en la rama, obj es distinto de nil
		if obj != nil {
			//Se define una nueva raíz del árbol
			arbol.Raiz = &RamaB{Primero: nil, Hoja: true, Contador: 0}
			arbol.Raiz.InsertarNodo(obj)
			arbol.Raiz.Hoja = false
		}
	}
}

//Función para determinar en qué rama debe insertarse el nuevo nodo
func (arbol *ArbolB) insertarRama(nodo *NodoB, rama *RamaB) *NodoB {
	/*
	IF -> Si la rama donde se va a insertar el nodo es una hoja
	ELSE -> Si la rama donde se va a insertar el nodo NO es una hoja
	*/
	if rama.Hoja {
		rama.InsertarNodo(nodo)
		/*
		IF -> Si el contador de la rama es igual al orden del árbol (es decir, si la rama tiene 3 nodos insertados en ella), se hace una división en la rama. Se retorna el nodo raiz de dicha división
		ELSE -> El contador de la rama no es igual al orden del árbol
		*/
		if rama.Contador == arbol.Orden {
			return arbol.dividir(rama)
		} else {
			return nil
		}
	} else {
		temp := rama.Primero
		for temp != nil {
			/*
			IF -> Si el curso de temp (el que recorre el árbol) es igual al curso del nodo que se desea insertar, se retorna nil (no se inserta el nodo porque ya hay uno con el mismo dato)
			ELSE IF 1 -> Si el curso del nodo a insertar es menor al del temporal, se recorre el subárbol izquierdo
			ELSE IF 2 -> Si se llega al último nodo de la rama, se recorre el subárbol derecho
			*/
			if nodo.Tutor.Curso == temp.Tutor.Curso {
				return nil
			} else if nodo.Tutor.Curso < temp.Tutor.Curso {
				obj := arbol.insertarRama(nodo, temp.Izquierdo)
				//Si es distinto de nil, es porque se realizó una división
				if obj != nil {
					rama.InsertarNodo(obj)
					if rama.Contador == arbol.Orden {
						return arbol.dividir(rama)
					}
				}
				return nil
			} else if temp.Siguiente == nil {
				obj := arbol.insertarRama(nodo, temp.Derecho)
				//Si es distinto de nil, es porque se realizó una división
				if obj != nil {
					rama.InsertarNodo(obj)
					if rama.Contador == arbol.Orden {
						return arbol.dividir(rama)
					}
				}
				return nil
			}
			temp = temp.Siguiente
		}
	}
	return nil
}

//Función para realizar la división en la rama que ha superado la cantidad de elemntos permitida según el orden del árbol B
func (arbol *ArbolB) dividir(rama *RamaB) *NodoB {
	tutor := &Tutor{Carnet: 0, Nombre: "", Curso: "", Password: ""}
	//Nodo temporal, que contendrá a la raiz de la división
	val := &NodoB{Tutor: tutor}
	//Auxiliar para recorrer la lista (la rama)
	aux := rama.Primero
	//Rama derecha e izquierda auxiliares para la división
	rderecha := &RamaB{Primero: nil, Contador: 0, Hoja: true}
	rizquierda := &RamaB{Primero: nil, Contador: 0, Hoja: true}
	//Variable para identificar el nodo que se está manejando
	contador := 0
	//Recorrer la rama
	for aux != nil {
		contador++
		/*
		IF -> Si se está en el primer elemento de la rama
		ELSE IF -> Si se está en el segundo elemento de la rama
		ELSE -> Si se está en el último elemento de la rama
		*/
		if contador < 2 {
			//Copia del primer nodo de la rama
			temp := &NodoB{Tutor: aux.Tutor}
			temp.Izquierdo = aux.Izquierdo
			if contador == 1 {
				temp.Derecho = aux.Siguiente.Izquierdo
			}
			//Verificar si es un nodo hoja o un nodo raiz
			if temp.Derecho != nil && temp.Izquierdo != nil {
				rizquierda.Hoja = false
			}
			rizquierda.InsertarNodo(temp)
		} else if contador == 2 {
			//Copia del nodo medio
			val.Tutor = aux.Tutor
		} else {
			//Copia del último nodo de la rama
			temp := &NodoB{Tutor: aux.Tutor}
			temp.Izquierdo = aux.Izquierdo
			temp.Derecho = aux.Derecho
			//Verificar si es un nodo hoja o un nodo raiz
			if temp.Derecho != nil && temp.Izquierdo != nil {
				rderecha.Hoja = false
			}
			rderecha.InsertarNodo(temp)
		}
		aux = aux.Siguiente
	}
	nuevo := &NodoB{Tutor: val.Tutor}
	nuevo.Derecho = rderecha
	nuevo.Izquierdo = rizquierda
	/*
	rama: ( 15 | 20 | 25 )
	resultado:		20 -> nuevo
				15		25
	*/
	return nuevo
}

//Función para buscar un tutor, con carnet recibido como parámetro, en el árbol B
func (arbol *ArbolB) BuscarTutor(carnet string, listaSimple *ListaSimple) {
	carnet_tutor, _ := strconv.Atoi(carnet)
	arbol.buscarEnArbol(arbol.Raiz.Primero, carnet_tutor, listaSimple)
	if listaSimple.Longitud > 0 {
		fmt.Println("Se encontró el tutor", listaSimple.Longitud)
	} else {
		fmt.Println("No se encontró")
	}
}

//Función donde se recorre el árbol B en busca de un tutor con número de carnet coincidente con el enviado como parámetro; si se encuentra, se inserta este tutor en la listaSimple
func (arbol *ArbolB) buscarEnArbol(raiz *NodoB, carnet_tutor int, listaSimple *ListaSimple) {
	if raiz != nil {
		aux := raiz
		//Recorrer el árbol B
		for aux != nil {
			/*
			IF 1 -> Buscar en la rama izquierda
			IF 2 -> Si coincide el carnet enviado y el carnet del nodo del árbol B
			IF 3 -> Buscar en la rama derecha
			*/
			if aux.Izquierdo != nil {
				arbol.buscarEnArbol(aux.Izquierdo.Primero, carnet_tutor, listaSimple)
			}
			if aux.Tutor.Carnet == carnet_tutor {
				listaSimple.InsertarTutor(aux)
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					arbol.buscarEnArbol(aux.Derecho.Primero, carnet_tutor, listaSimple)
				}
			}
			aux = aux.Siguiente
		}
	}
}

//Función para agregar un libro al arreglo de libros de un tutor de un tutor especifico
func (arbol *ArbolB) GuardarLibro(raiz *NodoB, nombre string, contenido string, carnet int) {
	if raiz != nil {
		aux := raiz
		//Se recorre el árbol B
		for aux != nil {
			//Recorrer rama izquierda
			if aux.Izquierdo != nil {
				arbol.GuardarLibro(aux.Izquierdo.Primero, nombre, contenido, carnet)
			}
			//Si el carnet del tutor del aux que recorre el árbol es igual al carnet enviado como parámetro
			if aux.Tutor.Carnet == carnet {
				//Se añade al arreglo de libros el libro a guardar
				raiz.Tutor.Libros = append(raiz.Tutor.Libros, &Libro{Nombre: nombre, Contenido: contenido, Estado: 1})
				fmt.Println("Registré el libro")
				return
			}
			//Recorrer rama derecha
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					arbol.GuardarLibro(aux.Derecho.Primero, nombre, contenido, carnet)
				}
			}
			aux = aux.Siguiente
		}
	}
}

//Función para agregar una publicacion al arreglo de publicaciones de un tutor especifico
func (arbol *ArbolB) GuardarPublicacion(raiz *NodoB, contenido string, carnet int) {
	if raiz != nil {
		aux := raiz
		//Se recorre el árbol B
		for aux != nil {
			//Recorrer rama izquierda
			if aux.Izquierdo != nil {
				arbol.GuardarPublicacion(aux.Izquierdo.Primero, contenido, carnet)
			}
			//Si el carnet del tutor del aux que recorre el árbol es igual al carnet enviado como parámetro
			if aux.Tutor.Carnet == carnet {
				//Se añade al arreglo de publicaciones la publicación a guardar
				raiz.Tutor.Publicaciones = append(raiz.Tutor.Publicaciones, &Publicacion{Contenido: contenido})
				fmt.Println("Registré la publicación")
				return
			}
			//Recorrer rama derecha
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					arbol.GuardarPublicacion(aux.Derecho.Primero, contenido, carnet)
				}
			}
			aux = aux.Siguiente
		}
	}
}

//Función para generar el reporte del árbol B en .dot y .jpg
func (arbol *ArbolB) ReporteArbolBTutores(nombre string) {
	cadena := ""
	nombre_archivo := "./" + nombre + ".dot"
	nombre_imagen := nombre + ".jpg"
	if arbol.Raiz != nil {
		cadena += "digraph arbol { \nnode[shape=record]\n"
		cadena += arbol.grafo(arbol.Raiz.Primero)
		cadena += arbol.conexionRamas(arbol.Raiz.Primero)
		cadena += "}"
	}
	generarArchivos.CrearArchivo(nombre_archivo)
	generarArchivos.EscribirArchivo(cadena, nombre_archivo)
	generarArchivos.Ejecutar(nombre_imagen, nombre_archivo)
}

func (arbol *ArbolB) grafo(rama *NodoB) string {
	dot := ""
	if rama != nil {
		dot += arbol.grafoRamas(rama)
		aux := rama
		for aux != nil {
			if aux.Izquierdo != nil {
				dot += arbol.grafo(aux.Izquierdo.Primero)
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					dot += arbol.grafo(aux.Derecho.Primero)
				}
			}
			aux = aux.Siguiente
		}
	}
	return dot
}

func (arbol *ArbolB) grafoRamas(rama *NodoB) string {
	dot := ""
	if rama != nil {
		aux := rama
		dot = dot + "R" + rama.Tutor.Curso + "[label=\""
		r := 1
		for aux != nil {
			if aux.Izquierdo != nil {
				dot = dot + "<C" + strconv.Itoa(r) + ">|"
				r++
			}
			if aux.Siguiente != nil {
				dot = dot + aux.Tutor.Curso + "|"
			} else {
				dot = dot + aux.Tutor.Curso
				if aux.Derecho != nil {
					dot = dot + "|<C" + strconv.Itoa(r) + ">"
				}
			}
			aux = aux.Siguiente
		}
		dot = dot + "\"];\n"
	}
	return dot
}

func (arbol *ArbolB) conexionRamas(rama *NodoB) string {
	dot := ""
	if rama != nil {
		aux := rama
		actual := "R" + rama.Tutor.Curso
		r := 1
		for aux != nil {
			if aux.Izquierdo != nil {
				dot += actual + ":C" + strconv.Itoa(r) + " -> " + "R" + aux.Izquierdo.Primero.Tutor.Curso + ";\n"
				r++
				dot += arbol.conexionRamas(aux.Izquierdo.Primero)
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					dot += actual + ":C" + strconv.Itoa(r) + " -> " + "R" + aux.Derecho.Primero.Tutor.Curso + ";\n"
					r++
					dot += arbol.conexionRamas(aux.Derecho.Primero)
				}
			}
			aux = aux.Siguiente
		}
	}
	return dot
}