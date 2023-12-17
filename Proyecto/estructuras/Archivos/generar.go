package Archivos

import (
	"fmt"
	"os"
	"os/exec"
)

func CrearArchivo(nombre_archivo string) {
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

func EscribirArchivo(contenido string, nombre_archivo string) {
	var archivo, err = os.OpenFile(nombre_archivo, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer archivo.Close()
	_, err = archivo.WriteString(contenido)
	if err != nil {
		return
	}
	err = archivo.Sync()
	if err != nil {
		return
	}
	fmt.Println("Archivo guardado con éxito")
}

func Ejecutar(nombre_imagen string, archivo string) {
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tjpg", archivo).Output()
	mode := 0777
	_ = os.WriteFile(nombre_imagen, cmd, os.FileMode(mode))
}