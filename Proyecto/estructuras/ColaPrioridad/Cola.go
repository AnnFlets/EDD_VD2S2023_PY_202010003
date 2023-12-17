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

func (cola *Cola) LeerCSVTutores(ruta string){
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
		nota, _ := strconv.Atoi(linea[3])
		cola.Encolar(carnet_csv, linea[1], "0" + linea[2], nota)
	}
	fmt.Println("Tutores cargados con éxito")
}

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
		for aux.Siguiente != nil{
			if aux.Siguiente.Prioridad > nuevoNodo.Prioridad && (aux.Prioridad == nuevoNodo.Prioridad || aux.Prioridad < nuevoNodo.Prioridad){
				nuevoNodo.Siguiente = aux.Siguiente
				aux.Siguiente = nuevoNodo
				cola.Tamanio++
				return
			}else{
				aux = aux.Siguiente
			}
		}
		aux.Siguiente = nuevoNodo
		cola.Tamanio++
	}
}

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

func (cola *Cola) Descolar(){
	if cola.Tamanio == 0{
		fmt.Println("No hay tutores en la cola")
	}else{
		cola.Primero = cola.Primero.Siguiente
		cola.Tamanio--
	}
}