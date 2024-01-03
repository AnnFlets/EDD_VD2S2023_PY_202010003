package main

import (
	"Fase2/estructuras/arbolB"
	"Fase2/estructuras/grafo"
	"Fase2/estructuras/peticiones"
	"Fase2/estructuras/tablaHash"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var tablaAlumnos *tablaHash.TablaHash
var listaSimple *arbolB.ListaSimple
var arbolTutor *arbolB.ArbolB
var grafoCursos *grafo.Grafo

func main() {
	tablaAlumnos = &tablaHash.TablaHash{Tabla: make(map[int]tablaHash.NodoHash), Capacidad: 7, Utilizacion: 0}
	listaSimple = &arbolB.ListaSimple{Inicio: nil, Longitud: 0}
	arbolTutor = &arbolB.ArbolB{Raiz: nil, Orden: 3}
	grafoCursos = &grafo.Grafo{Principal: &grafo.NodoGrafo{Curso: "ECYS"}}
	app := fiber.New()
	app.Use(cors.New())
	app.Post("/login", ValidarUsuario)
	app.Post("/registrar-alumno", RegistrarAlumno)
	app.Post("/registrar-tutor", RegistrarTutor)
	app.Post("/registrar-cursos", RegistrarCursos)
	app.Get("/tabla-alumnos", TablaAlumnos)
	app.Post("/registrar-libro", GuardarLibro)
	app.Listen(":4000")
}

func SHA256(cadena string) string {
	hexaString := ""
	h := sha256.New()
	h.Write([]byte(cadena))
	hexaString = hex.EncodeToString(h.Sum(nil))
	return hexaString
}

func ValidarUsuario(c *fiber.Ctx) error {
	var usuario peticiones.PeticionLogin
	listaSimple = &arbolB.ListaSimple{Inicio: nil, Longitud: 0}
	c.BodyParser(&usuario)
	if usuario.UserName == "ADMIN_202010003" {
		if usuario.Password == "Admin" {
			return c.JSON(&fiber.Map{
				"status":  200,
				"message": "Credenciales correctas",
				"rol":     1,
			})
		}
	} else {
		if usuario.Tutor {
			arbolTutor.BuscarTutor(usuario.UserName, listaSimple)
			if listaSimple.Longitud > 0 {
				if listaSimple.Inicio.Tutor.Tutor.Password == SHA256(usuario.Password) {
					return c.JSON(&fiber.Map{
						"status":  200,
						"message": "Credenciales correctas",
						"rol":     2,
					})
				}
			}
		} else {
			if tablaAlumnos.BuscarEstudiante(usuario.UserName, SHA256(usuario.Password)) {
				return c.JSON(&fiber.Map{
					"status":  200,
					"message": "Credenciales correctas",
					"rol":     3,
				})
			}
		}
	}
	return c.JSON(&fiber.Map{
		"status":  400,
		"message": "Credenciales incorrectas",
		"rol":     0,
	})
}

func RegistrarAlumno(c *fiber.Ctx) error {
	var alumno peticiones.PeticionRegistroAlumno
	c.BodyParser(&alumno)
	fmt.Println(alumno)
	tablaAlumnos.InsertarEstudiante(alumno.Carnet, alumno.Nombre, SHA256(alumno.Password), alumno.Cursos)
	return c.JSON(&fiber.Map{
		"status":  200,
		"Arreglo": tablaAlumnos.ConvertirArreglo(),
	})
}

func RegistrarTutor(c *fiber.Ctx) error {
	var tutor peticiones.PeticionRegistroTutor
	c.BodyParser(&tutor)
	arbolTutor.InsertarTutor(tutor.Carnet, tutor.Nombre, tutor.Curso, SHA256(tutor.Password))
	return c.JSON(&fiber.Map{
		"status": 200,
	})
}

func TablaAlumnos(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"status":  200,
		"Arreglo": tablaAlumnos.ConvertirArreglo(),
	})
}

func RegistrarCursos(c *fiber.Ctx) error {
	var cursito peticiones.PeticionCursos
	c.BodyParser(&cursito)
	fmt.Println(cursito)
	for _, curso := range cursito.Cursos {
		if len(curso.Post) > 0 {
			for j := 0; j < len(curso.Post); j++ {
				grafoCursos.InsertarCurso(curso.Codigo, curso.Post[j])
			}
		} else {
			grafoCursos.InsertarCurso("ECYS", curso.Codigo)
		}
	}
	return c.JSON(&fiber.Map{
		"status": 200,
	})
}

func GuardarLibro(c *fiber.Ctx) error {
	var libro peticiones.PeticionLibro
	c.BodyParser(&libro)
	fmt.Println(libro)
	arbolTutor.GuardarLibro(arbolTutor.Raiz.Primero, libro.Nombre, libro.Contenido, libro.Carnet)
	return c.JSON(&fiber.Map{
		"status": 200,
	})
}