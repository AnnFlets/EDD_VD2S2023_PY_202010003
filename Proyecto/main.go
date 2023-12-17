package main

import (
	"Proyecto/estructuras/ArbolAVL"
	"Proyecto/estructuras/ColaPrioridad"
	"Proyecto/estructuras/ListaCDE"
	"Proyecto/estructuras/ListaDE"
	"Proyecto/estructuras/MatrizDispersa"
	"fmt"
	"strconv"
)

var lista_doble *ListaDE.ListaDoble = &ListaDE.ListaDoble{Inicio: nil, Tamanio: 0}
var cola_prioridad *ColaPrioridad.Cola = &ColaPrioridad.Cola{Primero: nil, Tamanio: 0}
var lista_circular *ListaCDE.ListaCircular = &ListaCDE.ListaCircular{Inicio: nil, Tamanio: 0}
var matriz_dispersa *MatrizDispersa.Matriz = &MatrizDispersa.Matriz{Raiz: &MatrizDispersa.NodoMatriz{PosX: -1, PosY: -1, Dato: &MatrizDispersa.Dato{Carnet_Tutor: 0, Carnet_Estudiante: 0, Curso: "RAIZ"}}, Cantidad_Estudiantes: 0, Cantidad_Tutores: 0}
var arbol_avl *ArbolAVL.Arbol = &ArbolAVL.Arbol{Raiz: nil}
var estudiante_logeado string = ""

func main() {
	opcion := 0
	menu_principal := true
	for menu_principal {
		fmt.Println("========= MENÚ =========")
		fmt.Println("| 1. Inicio de sesión  |")
		fmt.Println("| 2. Salir             |")
		fmt.Println("========================")
		fmt.Println("Ingrese la opción a realizar:")
		fmt.Scanln(&opcion)
		if (opcion == 1){
			Login()
		}else if (opcion == 2){
			menu_principal = false
		}else{
			fmt.Println("[ERROR]: Opción inválida")
		}
	}
}

func Login(){
	usuario := ""
	contrasena := ""
	fmt.Println("========== LOGIN ==========")
	fmt.Print("Usuario: ")
	fmt.Scanln(&usuario)
	fmt.Print("Contraseña: ")
	fmt.Scanln(&contrasena)
	if usuario == "ADMIN_202010003" && contrasena == "Admin"{
		MenuAdministrador()
	}else if lista_doble.BuscarEstudiante(usuario, contrasena){
		estudiante_logeado = usuario
		MenuEstudiante()
	}else{
		fmt.Println("[ERROR]: Credenciales inválidas")
	}
}

func MenuAdministrador(){
	opcion := 0
	menu_admin := true
	for menu_admin {
		fmt.Println("========= MENÚ ADMINISTRADOR ========")
		fmt.Println("| 1. Carga de estudiantes tutores   |")
		fmt.Println("| 2. Carga de estudiantes           |")
		fmt.Println("| 3. Cargar cursos al sistema       |")
		fmt.Println("| 4. Control de estudiantes tutores |")
		fmt.Println("| 5. Reportes                       |")
		fmt.Println("| 6. Salir                          |")
		fmt.Println("=====================================")
		fmt.Println("Ingrese la opción a realizar:")
		fmt.Scanln(&opcion)
		switch opcion{
		case 1:
			CargarTutores()
		case 2:
			CargarEstudiantes()
		case 3:
			CargarCursos()
		case 4:
			ControlTutores()
		case 5:
			Reportes()
		case 6:
			menu_admin = false
		default:
			fmt.Println("[ERROR]: Opción inválida")
		}
	}
}

func CargarTutores() {
	ruta := ""
	fmt.Println("- CARGAR TUTORES -")
	fmt.Println("Ingrese el nombre del archivo: ")
	fmt.Scanln(&ruta)
	cola_prioridad.LeerCSVTutores(ruta)
}

func CargarEstudiantes() {
	ruta := ""
	fmt.Println("- CARGAR ESTUDIANTES -")
	fmt.Println("Ingrese el nombre del archivo: ")
	fmt.Scanln(&ruta)
	lista_doble.LeerCSVEstudiantes(ruta)
}

func CargarCursos() {
	ruta := ""
	fmt.Println("- CARGAR CURSOS -")
	fmt.Println("Ingrese el nombre del archivo: ")
	fmt.Scanln(&ruta)
	arbol_avl.LeerJsonCursos(ruta)
}

func ControlTutores() {
	opcion := 0
	menu_control_tutores := true
	for menu_control_tutores {
		fmt.Println("- CONTROL TUTORES -")
		cola_prioridad.Mostrar_Primero_Cola()
		fmt.Println("================================")
		fmt.Println("1. Aceptar")
		fmt.Println("2. Rechazar")
		fmt.Println("3. Salir")
		fmt.Println("Ingrese la opción a realizar:")
		fmt.Scanln(&opcion)
		if opcion == 1 {
			tutor_curso_buscar := lista_circular.BuscarTutor(cola_prioridad.Primero.Tutor.Curso)
			if tutor_curso_buscar != nil{
				if  cola_prioridad.Primero.Tutor.Nota > tutor_curso_buscar.Tutor.Nota{
					tutor_curso_buscar.Tutor.Carnet = cola_prioridad.Primero.Tutor.Carnet
					tutor_curso_buscar.Tutor.Nombre = cola_prioridad.Primero.Tutor.Nombre
					tutor_curso_buscar.Tutor.Curso = cola_prioridad.Primero.Tutor.Curso
					tutor_curso_buscar.Tutor.Nota = cola_prioridad.Primero.Tutor.Nota
					fmt.Println("Se sustituyó al tutor del curso actual")
				}
			}else{
				lista_circular.InsertarTutor(cola_prioridad.Primero.Tutor.Carnet, cola_prioridad.Primero.Tutor.Nombre, cola_prioridad.Primero.Tutor.Curso, cola_prioridad.Primero.Tutor.Nota)
				fmt.Println("Se registró tutor con éxito")
			}
			cola_prioridad.Descolar()
		} else if opcion == 2 {
			cola_prioridad.Descolar()
		} else if opcion == 3 {
			menu_control_tutores = false
		} else {
			fmt.Println("[ERROR]: Opción inválida")
		}
	}
}

func Reportes() {
	opcion := 0
	menu_reportes := true
	for menu_reportes {
		fmt.Println("============= REPORTES ============")
		fmt.Println("| 1. Reporte de alumnos           |")
		fmt.Println("| 2. Reporte de tutores aceptados |")
		fmt.Println("| 3. Reporte de asignacion        |")
		fmt.Println("| 4. Reporte de cursos            |")
		fmt.Println("| 5. Salir                        |")
		fmt.Println("===================================")
		fmt.Println("Ingrese la opción a realizar:")
		fmt.Scanln(&opcion)
		switch opcion{
		case 1:
			lista_doble.ReporteEstudiantes()
		case 2:
			lista_circular.ReporteTutores()
		case 3:
			matriz_dispersa.ReporteAsignaciones()
		case 4:
			arbol_avl.ReporteCursos()
		case 5:
			menu_reportes = false
		default:
			fmt.Println("[ERROR]: Opción inválida")
		}
	}
}

func MenuEstudiante(){
	opcion := 0
	menu_estudiante := true
	for menu_estudiante {
		fmt.Println("Bienvenido:", estudiante_logeado)
		fmt.Println("======= MENÚ ESTUDIANTE ======")
		fmt.Println("| 1. Ver tutores disponibles |")
		fmt.Println("| 2. Asignarse un tutor      |")
		fmt.Println("| 3. Salir                   |")
		fmt.Println("==============================")
		fmt.Println("Ingrese la opción a realizar:")
		fmt.Scanln(&opcion)
		switch opcion{
		case 1:
			fmt.Println("- TUTORES DISPONIBLES -")
			if lista_circular.Tamanio == 0{
				fmt.Println("No hay tutores disponibles")
			}else{
				lista_circular.MostrarTutores()
			}
		case 2:
			AsignarCurso()
		case 3:
			menu_estudiante = false
		default:
			fmt.Println("[ERROR]: Opción inválida")
		}
	}
}

func AsignarCurso() {
	curso_asignar := ""
	menu_asignar_curso := true
	for menu_asignar_curso {
		fmt.Println("Escriba el código del curso a asignar:")
		fmt.Scanln(&curso_asignar)
		if arbol_avl.BuscarCurso(curso_asignar) {
			if lista_circular.BuscarCurso(curso_asignar) {
				tutor_curso := lista_circular.BuscarTutor(curso_asignar)
				estudiante_curso, err := strconv.Atoi(estudiante_logeado)
				if err != nil{
					break
				}
				matriz_dispersa.InsertarElemento(estudiante_curso, tutor_curso.Tutor.Carnet, curso_asignar)
				fmt.Println("Asignación exitosa")
				break
			} else {
				fmt.Println("[ERROR]: No hay tutores disponibles para el curso ingresado")
				menu_asignar_curso = false
			}
		} else {
			fmt.Println("[ERROR]: El curso no existe en el sistema")
			menu_asignar_curso = false
		}
	}
}