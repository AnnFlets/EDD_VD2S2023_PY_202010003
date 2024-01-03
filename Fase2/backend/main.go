package main

import (
	"Fase2/estructuras/arbolB"
	"Fase2/estructuras/arbolMerkle"
	"Fase2/estructuras/grafo"
	"Fase2/estructuras/peticiones"
	"Fase2/estructuras/tablaHash"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var tablaAlumnos *tablaHash.TablaHash
var listaSimple *arbolB.ListaSimple
var arbolTutor *arbolB.ArbolB
var grafoCursos *grafo.Grafo
var arbolLibros *arbolMerkle.ArbolMerkle

func main() {
	tablaAlumnos = &tablaHash.TablaHash{Tabla: make(map[int]tablaHash.NodoHash), Capacidad: 7, Utilizacion: 0}
	listaSimple = &arbolB.ListaSimple{Inicio: nil, Longitud: 0}
	arbolTutor = &arbolB.ArbolB{Raiz: nil, Orden: 3}
	grafoCursos = &grafo.Grafo{Principal: &grafo.NodoGrafo{Valor: "ECYS"}}
	arbolLibros = &arbolMerkle.ArbolMerkle{Raiz: nil, BloqueDatos: nil, CantidadBloques: 0}
	app := fiber.New()
	app.Use(cors.New())
	app.Post("/login", Validar)
	app.Post("/registrar-alumno", RegistrarAlumno)
	app.Post("/registrar-tutor", RegistrarTutor)
	app.Post("/registrar-cursos", RegistrarCursos)
	app.Get("/tabla-alumnos", TablaAlumnos)
	app.Post("/registrar-libro", GuardarLibro)
	app.Get("/enviar-libros-admin", ObtenerLibrosAdmin)
	app.Post("/registrar-publicacion", GuardarPublicacion)
	app.Post("/registrar-log", RegistrarDecision)
	app.Post("/obtener-clases", CursosAlumnos)
	app.Get("/obtener-libros-alumno", ObtenerLibrosAlumno)
	app.Get("/obtener-publi-alumno", ObtenerPublicacionesAlumno)
	app.Get("/finalizar-libros", FinalizarLibros)

	app.Get("/reporte-arbol", ReporteArbolB)
	app.Get("/reporte-grafo", ReporteGrafo)
	app.Get("/reporte-bloque", ReporteMerkle)
	app.Listen(":4000")
}

func Validar(c *fiber.Ctx) error {
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
				if listaSimple.Inicio.Tutor.Valor.Password == SHA256(usuario.Password) {
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

func SHA256(cadena string) string {
	hexaString := ""
	h := sha256.New()
	h.Write([]byte(cadena))
	hexaString = hex.EncodeToString(h.Sum(nil))
	return hexaString
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
	fmt.Println(libro.Nombre)
	arbolTutor.GuardarLibro(arbolTutor.Raiz.Primero, libro.Nombre, libro.Contenido, libro.Carnet)
	return c.JSON(&fiber.Map{
		"status": 200,
	})
}

func ObtenerLibrosAdmin(c *fiber.Ctx) error {
	listatemp := &arbolB.ListaSimple{Inicio: nil, Longitud: 0}
	var libros []arbolB.Libro
	arbolTutor.VerLibroAdmin(arbolTutor.Raiz.Primero, listatemp)
	if listatemp.Longitud > 0 {
		aux := listatemp.Inicio
		for aux != nil {
			for i := 0; i < len(aux.Tutor.Valor.Libros); i++ {
				if aux.Tutor.Valor.Libros[i].Estado == 1 {
					libros = append(libros, *aux.Tutor.Valor.Libros[i])
				}
			}
			aux = aux.Siguiente
		}
	}
	if len(libros) > 0 {
		return c.JSON(&fiber.Map{
			"status":  200,
			"Arreglo": libros,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 500,
	})
}

func GuardarPublicacion(c *fiber.Ctx) error {
	var publicacion peticiones.PeticionPublicacion
	c.BodyParser(&publicacion)
	arbolTutor.GuardarPublicacion(arbolTutor.Raiz.Primero, publicacion.Contenido, publicacion.Carnet)
	return c.JSON(&fiber.Map{
		"status": 200,
	})
}

func RegistrarDecision(c *fiber.Ctx) error {
	var accion peticiones.PeticionDecision
	c.BodyParser(&accion)
	arbolLibros.AgregarBloque(accion.Accion, accion.Nombre, accion.Tutor)
	if accion.Accion == "Aceptado" {
		arbolTutor.ActualizarLibro(arbolTutor.Raiz.Primero, accion.Nombre, accion.Curso, 2)
	} else if accion.Accion == "Rechazado" {
		arbolTutor.ActualizarLibro(arbolTutor.Raiz.Primero, accion.Nombre, accion.Curso, 3)
	}
	return c.JSON(&fiber.Map{
		"status": 200,
	})
}

func CursosAlumnos(c *fiber.Ctx) error {
	var alumno peticiones.PeticionAlumnoSesion
	c.BodyParser(&alumno)
	busqueda := tablaAlumnos.BuscarSesion(alumno.Carnet)
	if busqueda != nil {
		return c.JSON(&fiber.Map{
			"status":  200,
			"Arreglo": busqueda.Cursos,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 500,
	})
}

func ReporteArbolB(c *fiber.Ctx) error {
	var imagen peticiones.RespuestaImagen = peticiones.RespuestaImagen{Nombre: "Reporte/ArbolB.jpg"}
	arbolTutor.ReporteArbolBTutores("ArbolB")
	/*INICIO*/
	imageBytes, err := os.ReadFile(imagen.Nombre)
	if err != nil {
		return c.JSON(&fiber.Map{
			"status": 404,
		})
	}
	// Codifica los bytes de la imagen en base64
	imagen.Imagenbase64 = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(imageBytes)
	return c.JSON(&fiber.Map{
		"status": 200,
		"imagen": imagen,
	})
}

func ReporteGrafo(c *fiber.Ctx) error {
	var imagen peticiones.RespuestaImagen = peticiones.RespuestaImagen{Nombre: "Reporte/Grafo.jpg"}
	grafoCursos.ReporteGrafoCursos("Grafo")
	/*INICIO*/
	imageBytes, err := os.ReadFile(imagen.Nombre)
	if err != nil {
		return c.JSON(&fiber.Map{
			"status": 404,
		})
	}
	// Codifica los bytes de la imagen en base64
	imagen.Imagenbase64 = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(imageBytes)
	return c.JSON(&fiber.Map{
		"status": 200,
		"imagen": imagen,
	})
}

func ReporteMerkle(c *fiber.Ctx) error {
	var imagen peticiones.RespuestaImagen = peticiones.RespuestaImagen{Nombre: "Reporte/arbolMerkle.jpg"}
	arbolLibros.ReporteMerkleLibros("arbolMerkle")
	/*INICIO*/
	imageBytes, err := os.ReadFile(imagen.Nombre)
	if err != nil {
		return c.JSON(&fiber.Map{
			"status": 404,
		})
	}
	// Codifica los bytes de la imagen en base64
	imagen.Imagenbase64 = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(imageBytes)
	return c.JSON(&fiber.Map{
		"status": 200,
		"imagen": imagen,
	})
}

func ObtenerLibrosAlumno(c *fiber.Ctx) error {
	listatemp := &arbolB.ListaSimple{Inicio: nil, Longitud: 0}
	var libros []arbolB.Libro
	arbolTutor.VerLibroAdmin(arbolTutor.Raiz.Primero, listatemp)
	if listatemp.Longitud > 0 {
		aux := listatemp.Inicio
		for aux != nil {
			for i := 0; i < len(aux.Tutor.Valor.Libros); i++ {
				if aux.Tutor.Valor.Libros[i].Estado == 2 {
					libros = append(libros, *aux.Tutor.Valor.Libros[i])
				}
			}
			aux = aux.Siguiente
		}
	}
	if len(libros) > 0 {
		return c.JSON(&fiber.Map{
			"status":  200,
			"Arreglo": libros,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 500,
	})
}

func ObtenerPublicacionesAlumno(c *fiber.Ctx) error {
	listatemp := &arbolB.ListaSimple{Inicio: nil, Longitud: 0}
	var publi []arbolB.Publicacion
	arbolTutor.VerLibroAdmin(arbolTutor.Raiz.Primero, listatemp)
	if listatemp.Longitud > 0 {
		aux := listatemp.Inicio
		for aux != nil {
			for i := 0; i < len(aux.Tutor.Valor.Publicaciones); i++ {
				publi = append(publi, *aux.Tutor.Valor.Publicaciones[i])
			}
			aux = aux.Siguiente
		}
	}
	if len(publi) > 0 {
		return c.JSON(&fiber.Map{
			"status":  200,
			"Arreglo": publi,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 500,
	})
}

func FinalizarLibros(c *fiber.Ctx) error {
	arbolLibros.GenerarArbol()
	return c.JSON(&fiber.Map{
		"status": 200,
	})
}