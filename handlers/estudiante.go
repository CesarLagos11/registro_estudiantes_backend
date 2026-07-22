package handlers

import (
	"bufio"
	"fmt"

	"registro_es/estudiantes"
	"registro_es/models"
	"registro_es/storage"
)

func EjecutarMenu(reader *bufio.Reader) {
	repo := storage.NewArchivoEstudianteRepository(storage.ArchivoEstudiantes)
	gestor := estudiantes.NewGestor(repo)

	defer func() {
		if err := repo.Guardar(gestor.Listar()); err != nil {
			fmt.Printf("No se pudieron guardar los estudiantes: %v\n", err)
		}
		fmt.Println("Datos guardados en estudiantes.json")
	}()

	for {
		fmt.Println("\n=== Menú de estudiantes ===")
		fmt.Println("1. Registrar estudiante")
		fmt.Println("2. Consultar estudiante por ID")
		fmt.Println("3. Ver estudiantes")
		fmt.Println("4. Avanzar estudiante al siguiente grado")
		fmt.Println("5. Salir")
		opcion := LeerOpcion(reader, "Seleccione una opción: ")

		switch opcion {
		case 1:
			RegistrarEstudiante(reader, gestor)
		case 2:
			ConsultarEstudiante(gestor, reader)
		case 3:
			VerTodosLosEstudiantes(gestor, reader)
		case 4:
			AvanzarEstudiante(reader, gestor)
		case 5:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción inválida, ingrese un número del 1 al 5.")
		}
	}
}

func RegistrarEstudiante(reader *bufio.Reader, gestor *estudiantes.Gestor) {
	nombre := LeerLinea(reader, "Nombre: ")
	if nombre == "" {
		nombre = "No registrado"
	}

	edad := LeerEdad(reader, "Edad (deje en blanco para no registrar): ")
	correo := LeerCorreo(reader, "Correo (deje en blanco para no registrar): ")
	numCuenta := LeerNumCuenta(reader, "Número de cuenta (deje en blanco para no registrar): ")
	grado := LeerGrado(reader, "Grado (1-9, deje en blanco para 1): ")
	creditos := LeerCreditos(reader, "Créditos acumulados (deje en blanco para 0): ")

	estudiante := gestor.Registrar(models.Estudiante{
		Nombre:    nombre,
		Edad:      edad,
		Correo:    correo,
		NumCuenta: numCuenta,
		Grado:     grado,
		Creditos:  creditos,
	})
	fmt.Println("\nEstudiante registrado correctamente.")
	MostrarEstudiante(estudiante)

	for {
		fmt.Println("1. Agregar otro estudiante")
		fmt.Println("2. Volver al menú")
		opcion := LeerOpcion(reader, "Seleccione una opción: ")
		switch opcion {
		case 1:
			return
		case 2:
			return
		default:
			fmt.Println("Opción inválida")
		}
	}
}

func VerTodosLosEstudiantes(gestor *estudiantes.Gestor, reader *bufio.Reader) {
	fmt.Println("\n=== Todos los estudiantes registrados ===")
	estudiantes := gestor.Listar()
	if len(estudiantes) == 0 {
		fmt.Println("No hay estudiantes registrados.")
	} else {
		for _, estudiante := range estudiantes {
			MostrarEstudiante(estudiante)
		}
	}

	fmt.Println("Para volver al menú")
	LeerLinea(reader, "Presione Enter para continuar: ")
}

func ConsultarEstudiante(gestor *estudiantes.Gestor, reader *bufio.Reader) {
	id := LeerOpcion(reader, "ID del estudiante a consultar: ")
	if estudiante, ok := gestor.BuscarPorID(id); ok {
		MostrarEstudiante(estudiante)
		fmt.Println("Estado de avance:")
		if estudiantes.PuedeAvanzar(estudiante) {
			fmt.Printf("Sí puede avanzar al grado %d con %d créditos.\n", estudiante.Grado+1, storage.CreditosRequeridos)
		} else {
			fmt.Printf("No cumple con los requisitos para avanzar. Requiere %d créditos y estar en un grado menor a 9.\n", storage.CreditosRequeridos)
		}
		fmt.Println("Para volver al menú")
		LeerLinea(reader, "Presione Enter para continuar: ")
		return
	}

	fmt.Println("No se encontró un estudiante con ese ID.")
	LeerLinea(reader, "Presione Enter para continuar: ")
}

func AvanzarEstudiante(reader *bufio.Reader, gestor *estudiantes.Gestor) {
	id := LeerOpcion(reader, "ID del estudiante que desea avanzar: ")
	if estudiante, ok, err := gestor.Avanzar(id); err != nil {
		fmt.Println(err)
		LeerLinea(reader, "Presione Enter para continuar: ")
		return
	} else if ok {
		fmt.Printf("Estudiante %s avanzó al grado %d.\n", estudiante.Nombre, estudiante.Grado)
		LeerLinea(reader, "Presione Enter para continuar: ")
		return
	}

	fmt.Println("No se encontró un estudiante con ese ID.")
	LeerLinea(reader, "Presione Enter para continuar: ")
}

func MostrarEstudiante(e models.Estudiante) {
	fmt.Printf("ID: %d\n", e.ID)
	fmt.Printf("Estudiante: %s\n", e.Nombre)
	if e.Edad != nil {
		fmt.Printf("Edad: %d\n", *e.Edad)
	} else {
		fmt.Printf("Edad: No registrado\n")
	}
	if e.Correo != nil {
		fmt.Printf("Correo: %s\n", *e.Correo)
	} else {
		fmt.Printf("Correo: No registrado\n")
	}
	if e.NumCuenta != nil {
		fmt.Printf("Número de cuenta: %s\n", *e.NumCuenta)
	} else {
		fmt.Printf("Numero de cuenta: No registrado\n")
	}
	fmt.Printf("Grado: %d\n", e.Grado)
	fmt.Printf("Créditos: %d\n", e.Creditos)
	fmt.Println("-------------------------")
}

func PuedeAvanzar(estudiante models.Estudiante) bool {
	return estudiantes.PuedeAvanzar(estudiante)
}
