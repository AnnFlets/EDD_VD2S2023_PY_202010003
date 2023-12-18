package ColaPrioridad

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Cola struct {
	Primero  *NodoCola
	Tamanio int
}

//Función para leer el archivo csv con los datos de los tutores y agregarlos a la cola
func (cola *Cola) LeerCSVTutores(ruta string){
	archivo, err := os.Open(ruta)
	if err != nil{
		fmt.Println("[ERROR]: No se pudo abrir el archivo")
		return
	}
	defer archivo.Close()
	contenido := csv.NewReader(archivo)
	//Especificar el delimitador del archivo
	contenido.Comma = ','
	encabezado := true
	for {
		//Lee una línea del archivo
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
		nota, _ := strconv.Atoi(linea[3])
		cola.Encolar(carnet_csv, linea[1], "0" + linea[2], nota)
	}
	fmt.Println("Tutores cargados con éxito")
}

//Función para definir la prioridad del tutor e insertarlo en la cola según corresponda
func (cola *Cola) Encolar(carnet int, nombre string, curso string, nota int){
	nuevoTutor := &Tutor{Carnet: carnet, Nombre: nombre, Curso: curso, Nota: nota}
	nuevoNodo := &NodoCola{Tutor: nuevoTutor, Siguiente: nil, Prioridad: 0}
	if (nota >= 90 && nota <= 100){
		nuevoNodo.Prioridad = 1
	}else if (nota >= 75 && nota <= 89){
		nuevoNodo.Prioridad = 2
	}else if (nota >= 65 && nota <= 74){
		nuevoNodo.Prioridad = 3
	}else if (nota >= 61 && nota <= 64){
		nuevoNodo.Prioridad = 4
	}else{
		return
	}
	if cola.Tamanio == 0{
		cola.Primero = nuevoNodo
		cola.Tamanio++
	}else{
		/*
		IF -> Si el valor de la prioridad del primero de la cola es mayor a la del tutor a insertar
		ELSE IF ->  Si el valor de la prioridad del primero de la cola es menor a la del tutor a insertar, pero el primero de la cola es el único elemento de la cola.
		*/
		if cola.Primero.Prioridad > nuevoNodo.Prioridad{
			nuevoNodo.Siguiente = cola.Primero
			cola.Primero = nuevoNodo
			cola.Tamanio++
			return
		}else if cola.Primero.Siguiente == nil{
			cola.Primero.Siguiente = nuevoNodo
			cola.Tamanio++
			return
		}
		aux := cola.Primero
		//Se ejecuta hasta llegar al último nodo de la cola
		for aux.Siguiente != nil{
			/*
			IF -> Si el valor de la prioridad del siguiente es mayor a la del tutor a insertar y el valor de la prioridad del anterior es igual o menor a la del tutor a insertar
			ELSE ->  Si el valor de la prioridad del siguiente es menor a la del tutor a insertar
			*/
			if aux.Siguiente.Prioridad > nuevoNodo.Prioridad && (aux.Prioridad == nuevoNodo.Prioridad || aux.Prioridad < nuevoNodo.Prioridad){
				nuevoNodo.Siguiente = aux.Siguiente
				aux.Siguiente = nuevoNodo
				cola.Tamanio++
				return
			}else{
				aux = aux.Siguiente
			}
		}
		//Si se llega al último nodo, se inserta al nuevo tutor como el siguiente de este (de último)
		aux.Siguiente = nuevoNodo
		cola.Tamanio++
	}
}

//Función para mostrar los datos del primer tutor de la cola y el carnet del siguiente
func (cola *Cola) Mostrar_Primero_Cola(){
	if cola.Tamanio == 0{
		fmt.Println("No hay tutores")
	}else{
		fmt.Println("Actual: ", cola.Primero.Tutor.Carnet)
		fmt.Println("Nombre: ", cola.Primero.Tutor.Nombre)
		fmt.Println("Curso: ", cola.Primero.Tutor.Curso)
		fmt.Println("Nota: ", cola.Primero.Tutor.Nota)
		fmt.Println("Prioridad: ", cola.Primero.Prioridad)
		if cola.Primero.Siguiente != nil{
			fmt.Println("Siguiente:", cola.Primero.Siguiente.Tutor.Carnet)
		}else{
			fmt.Println("Siguiente: No hay más tutores por evaluar")
		}
	}
}

//Función para sacar al primer tutor de la cola
func (cola *Cola) Descolar(){
	if cola.Tamanio == 0{
		fmt.Println("No hay tutores en la cola")
	}else{
		cola.Primero = cola.Primero.Siguiente
		cola.Tamanio--
	}
}