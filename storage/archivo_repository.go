package storage

import (
	"registro_es/models"
)

type ArchivoEstudianteRepository struct {
	ruta string
}

func NewArchivoEstudianteRepository(ruta string) *ArchivoEstudianteRepository {
	return &ArchivoEstudianteRepository{ruta: ruta}
}

func (r *ArchivoEstudianteRepository) Cargar() []models.Estudiante {
	return CargarEstudiantesDesdeArchivo(r.ruta)
}

func (r *ArchivoEstudianteRepository) Guardar(estudiantes []models.Estudiante) error {
	return GuardarEstudiantesEnArchivo(estudiantes, r.ruta)
}
