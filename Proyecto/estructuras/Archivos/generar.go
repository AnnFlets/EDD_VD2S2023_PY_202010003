package Archivos

import (
	"fmt"
	"os"
	"os/exec"
)

//Función encargada de crear un archivo .dot
func CrearArchivo(nombre_archivo string) {
	//Busca el archivo
	var _, err = os.Stat(nombre_archivo)
	if os.IsNotExist(err) {
		var archivo, err = os.Create(nombre_archivo)
		if err != nil {
			return
		}
		defer archivo.Close()
	}
	fmt.Println("Archivo generado con éxito")
}

//Función para escribir el contenido recibido en el archivo .dot
func EscribirArchivo(contenido string, nombre_archivo string) {
	//Abre el archivo creado con permisos de lectura y escritura
	var archivo, err = os.OpenFile(nombre_archivo, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer archivo.Close()
	//Escribe el contenido deseado en el archivo
	_, err = archivo.WriteString(contenido)
	if err != nil {
		return
	}
	//Sincroniza que lo escrito se refleje en el archivo
	err = archivo.Sync()
	if err != nil {
		return
	}
	fmt.Println("Archivo guardado con éxito")
}

//Función para generar la imagen .jpg de acuerdo al .dot
func Ejecutar(nombre_imagen string, archivo string) {
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tjpg", archivo).Output()
	mode := 0777
	_ = os.WriteFile(nombre_imagen, cmd, os.FileMode(mode))
}