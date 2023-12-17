package ListaDE

import (
	"Proyecto/estructuras/Archivos"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type ListaDoble struct {
	Inicio  *NodoListaDoble
	Tamanio int
}

func (lista *ListaDoble) LeerCSVEstudiantes(ruta string){
	archivo, err := os.Open(ruta)
	if err != nil{
		fmt.Println("[ERROR]: No se pudo abrir el archivo")
		return
	}
	defer archivo.Close()
	contenido := csv.NewReader(archivo)
	contenido.Comma = ','
	encabezado := true
	for {
		linea, err := contenido.Read()
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Println("[ERROR]: No se pudo leer la línea")
			continue
		}
		if encabezado{
			encabezado = false
			continue
		}
		carnet_csv, _ := strconv.Atoi(linea[0])
		lista.InsertarEstudiante(carnet_csv, linea[1])
	}
	fmt.Println("Estudiantes cargados con éxito")
}

func (lista *ListaDoble) InsertarEstudiante(carnet int, nombre string) {
	nuevoEstudiante := &Estudiante{Carnet: carnet, Nombre: nombre}
	nuevoNodo := &NodoListaDoble{Estudiante: nuevoEstudiante, Siguiente: nil, Anterior: nil}
	if lista.Tamanio == 0 {
		lista.Inicio = nuevoNodo
	} else {
		aux := lista.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		nuevoNodo.Anterior = aux
		aux.Siguiente = nuevoNodo
	}
	lista.Tamanio++
}

func (lista *ListaDoble) BuscarEstudiante(carnet string, contrasena string) bool{
	if lista.Tamanio == 0{
		return false
	}else{
		aux := lista.Inicio
		for aux != nil{
			if (strconv.Itoa(aux.Estudiante.Carnet) == carnet && carnet == contrasena){
				return true
			}
			aux = aux.Siguiente
		}
	}
	return false
}

func (lista *ListaDoble) ReporteEstudiantes(){
	nombre_archivo := "./listadoble.dot"
	nombre_imagen := "./listadoble.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull1[label=\"null\"];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := lista.Inicio
	contador := 0
	texto += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < lista.Tamanio; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + "Nombre: " + aux.Estudiante.Nombre + "\\n" + "Carnet: " + strconv.Itoa(aux.Estudiante.Carnet) + "\"];\n"
		aux = aux.Siguiente
	}
	for i := 0; i < (lista.Tamanio - 1); i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	Archivos.CrearArchivo(nombre_archivo)
	Archivos.EscribirArchivo(texto, nombre_archivo)
	Archivos.Ejecutar(nombre_imagen, nombre_archivo)
}