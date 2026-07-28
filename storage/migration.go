package storage

import (
	"errors"
	"fmt"

	"registro_es/models"

	"gorm.io/gorm"
)

func MigrarDatosDesdeJSON(db *gorm.DB) {
	estudiantesJSON := CargarEstudiantesDesdeArchivo(ArchivoEstudiantes)
	if len(estudiantesJSON) == 0 {
		return
	}

	for _, estudiante := range estudiantesJSON {
		var existente models.Estudiante
		if err := db.First(&existente, estudiante.ID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := db.Create(&estudiante).Error; err != nil {
					fmt.Printf("No se pudo migrar estudiante ID %d: %v\n", estudiante.ID, err)
				}
			}
		}
	}
}
