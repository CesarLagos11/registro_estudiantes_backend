package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"registro_es/models"
)

const (
	ArchivoEstudiantes = "estudiantes.json"
	CreditosRequeridos = 7
	GradoMinimo        = 1
	GradoMaximo        = 9
)

func GuardarEstudiantesEnArchivo(estudiantes []models.Estudiante, ruta string) error {
	contenido, err := json.MarshalIndent(estudiantes, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(ruta, contenido, 0o644)
}

func CargarEstudiantesDesdeArchivo(ruta string) []models.Estudiante {
	contenido, err := os.ReadFile(ruta)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Estudiante{}
		}
		fmt.Printf("No se pudo leer el archivo %s: %v\n", ruta, err)
		return []models.Estudiante{}
	}

	var estudiantes []models.Estudiante
	if err := json.Unmarshal(contenido, &estudiantes); err != nil {
		fmt.Printf("No se pudo leer la información de %s: %v\n", ruta, err)
		return []models.Estudiante{}
	}
	return estudiantes
}
