package tablaHash

import "strconv"

/*
Estructura:
	Tabla = {
		3: nodoHash3,
		4: nodoHash4,
		...
	}
*/

type TablaHash struct {
	//Tipo map, donde la llave será int y el dato a almacenar NodoHash
	Tabla map[int]NodoHash
	//Cantidad de elementos que puede almacenar la tabla
	Capacidad int
	//Cantidad de elementos almacenados en la tabla
	Utilizacion int
}

//Función para insertar un estudiante en la tabla hash
func (tabla *TablaHash) InsertarEstudiante(carnet int, nombre string, password string, cursos []string) {
	indice := tabla.calcularIndice(carnet)
	nuevoNodo := &NodoHash{Llave: indice, Estudiante: &Estudiante{Carnet: carnet, Nombre: nombre, Password: password, Cursos: cursos}}
	if tabla.Capacidad > indice {
		/*
		IF -> NO existe algún estudiante en el índice (posición) calculado
		ELSE -> Si hay colisión (ya existe un elemento en el índice calculado)
		*/
		if _, existe := tabla.Tabla[indice]; !existe{
			tabla.Tabla[indice] = *nuevoNodo
			tabla.Utilizacion++
			tabla.capacidadTabla()
		}else{
			intento := 1
			indice = tabla.recalcularIndice(carnet, intento)
			for {
				/*
				IF -> Si hay colisión (existe un estudiante en el índice calculado)
				ELSE -> NO existe algún estudiante en el índice (posición) calculado
				*/
				if _, existe := tabla.Tabla[indice]; existe{
					intento++
					indice = tabla.recalcularIndice(carnet, intento)
				}else{
					nuevoNodo.Llave = indice
					tabla.Tabla[indice] = *nuevoNodo
					tabla.Utilizacion++
					tabla.capacidadTabla()
					break
				}
			}
		}
	}
}

//Función para calcular y retornar el índice (posición), de acuerdo al número de carnet, en la que debe almacenarse un elemento en la tabla hash.
func (tabla *TablaHash) calcularIndice(carnet int) int {
	var numeros []int
	//Convertir el carnet en un arreglo de digitos
	for {
		if carnet > 0 {
			//Tomar el último dígito del carnet
			digito := carnet % 10
			//Agregarlo al arreglo de números
			numeros = append([]int{digito}, numeros...)
			//Quitar el último dígito del carnet (que ya se agregó al arreglo)
			carnet = carnet / 10
		} else {
			break
		}
	}
	var numeros_ascii []rune
	//Recorrer el arreglo de dígitos y convertir cada uno a su código ASCII
	for _, numero := range numeros {
		valor_ascii := rune(numero + 48)
		numeros_ascii = append(numeros_ascii, valor_ascii)
	}
	suma_ascii := 0
	//Recorrer el arreglo con los valores ASCII, sumando cada uno de los valores
	for _, numero_ascii := range numeros_ascii {
		suma_ascii += int(numero_ascii)
	}
	indice := suma_ascii % tabla.Capacidad
	return indice
}

//Función para verificar aspectos relacionados a la capacidad de la tabla hash y ver si es factible aumentar esta o no.
func (tabla *TablaHash) capacidadTabla() {
	porcentaje_utilizacion := float64(tabla.Capacidad) * 0.7
	//Si la cantidad de elementos almacenados en la tabla supera el porcentaje de utilización
	if tabla.Utilizacion > int(porcentaje_utilizacion) {
		capacidad_anterior := tabla.Capacidad
		tabla.Capacidad = tabla.nuevaCapacidad()
		tabla.Utilizacion = 0
		tabla.insertarEnNuevaTabla(capacidad_anterior)
	}
}

//Función para encontrar y retornar la nueva capacidad de la tabla hash, basado en la serie Fibonacci.
func (tabla *TablaHash) nuevaCapacidad() int {
	contador := 0
	anterior, siguiente := 0, 1
	for contador < 100 {
		contador++
		if anterior > tabla.Capacidad {
			return anterior
		}
		anterior, siguiente = siguiente, anterior + siguiente
	}
	return anterior
}

//Función para insertar los datos de la tabla hash anterior a la nueva tabla hash con mayor capacidad
func (tabla *TablaHash) insertarEnNuevaTabla(capacidad_anterior int) {
	//Tabla hash anterior
	aux_tabla := tabla.Tabla
	//Nueva tabla hash
	tabla.Tabla = make(map[int]NodoHash)
	//Se recorre la tabla hash anterior
	for i := 0; i < capacidad_anterior; i++ {
		//Si existe un estudiante en la posición i, se insertan los datos de este en la nueva tabla hash
		if usuario, existe := aux_tabla[i]; existe {
			tabla.InsertarEstudiante(usuario.Estudiante.Carnet, usuario.Estudiante.Nombre, usuario.Estudiante.Password, usuario.Estudiante.Cursos)
		}
	}
}

//Función que retorna el nuevo índice calculado por direccionamiento abierto cuadrático
func (tabla *TablaHash) recalcularIndice(carnet int, intento int) int {
	nuevo_indice := tabla.calcularIndice(carnet) + (intento * intento)
	return tabla.nuevoIndice(nuevo_indice)
}

//Función que retorna el nuevo índice, verificando que este no sobrepase la capacidad de la tabla y en dicho caso, recorrer nuevamente la tabla con los saltos restantes
func (tabla *TablaHash) nuevoIndice(nuevo_indice int) int {
	nueva_pos := 0
	if nuevo_indice < tabla.Capacidad {
		nueva_pos = nuevo_indice
	} else {
		nueva_pos = nuevo_indice - tabla.Capacidad
		nueva_pos = tabla.nuevoIndice(nueva_pos)
	}
	return nueva_pos
}

//Función para buscar un estudiante, por carnet y password, en la tabla hash. Retorna true si lo encuentra y en el caso contrario false.
func (tabla *TablaHash) BuscarEstudiante(carnet string, password string) bool {
	//Convertir el carnet a int
	carnet_estudiante, err := strconv.Atoi(carnet)
	if err != nil {
		return false
	}
	indice := tabla.calcularIndice(carnet_estudiante)
	if indice < tabla.Capacidad {
		/*
		IF -> Si existe un estudiante en el índice calculado
		*/
		if usuario, existe := tabla.Tabla[indice]; existe {
			if usuario.Estudiante.Carnet == carnet_estudiante {
				if usuario.Estudiante.Password == password {
					return true
				}else{
					return false
				}
			} else {
				intento := 1
				indice = tabla.recalcularIndice(carnet_estudiante, intento)
				for {
					/*
					IF -> Si existe un estudiante en el índice calculado
					*/
					if usuario, existe := tabla.Tabla[indice]; existe {
						if usuario.Estudiante.Carnet == carnet_estudiante {
							if usuario.Estudiante.Password == password {
								return true
							} else {
								return false
							}
						} else {
							intento++
							indice = tabla.recalcularIndice(carnet_estudiante, intento)
						}
					} else {
						return false
					}
				}
			}
		}
	}
	return false
}

//Función para retornar un arreglo con la información de los estudiantes almacenados en la tabla hash.
func (tabla *TablaHash) ConvertirArreglo() []NodoHash {
	var arrays []NodoHash
	if tabla.Utilizacion > 0 {
		for i := 0; i < tabla.Capacidad; i++ {
			if usuario, existe := tabla.Tabla[i]; existe {
				arrays = append(arrays, usuario)
			}
		}
	}
	return arrays
}

//Función que retorna al estudiante de la tabla hash con número de carnet igual al recibido como parámetro.
func (tabla *TablaHash) BuscarSesion(carnet string) *Estudiante {
	temp, err := strconv.Atoi(carnet)
	if err != nil {
		return nil
	}
	indice := tabla.calcularIndice(temp)
	if indice < tabla.Capacidad {
		if usuario, existe := tabla.Tabla[indice]; existe {
			if usuario.Estudiante.Carnet == temp {
				return usuario.Estudiante
			} else {
				contador := 1
				indice = tabla.recalcularIndice(temp, contador)
				for {
					if usuario, existe := tabla.Tabla[indice]; existe {
						if usuario.Estudiante.Carnet == temp {
							return usuario.Estudiante
						} else {
							contador++
							indice = tabla.recalcularIndice(temp, contador)
						}
					} else {
						return nil
					}
				}
			}
		}
	}
	return nil
}