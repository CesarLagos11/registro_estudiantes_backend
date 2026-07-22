package estudiantes

import (
	"errors"

	"registro_es/models"
	"registro_es/storage"
)

type Repositorio interface {
	Cargar() []models.Estudiante
	Guardar([]models.Estudiante) error
}

type Gestor struct {
	repositorio Repositorio
}

func NewGestor(repositorio Repositorio) *Gestor {
	return &Gestor{repositorio: repositorio}
}

func (g *Gestor) Listar() []models.Estudiante {
	return g.repositorio.Cargar()
}

func (g *Gestor) Registrar(estudiante models.Estudiante) models.Estudiante {
	estudiantes := g.repositorio.Cargar()
	id := siguienteID(estudiantes)
	estudiante.ID = id
	if estudiante.Nombre == "" {
		estudiante.Nombre = "No registrado"
	}
	if estudiante.Grado == 0 {
		estudiante.Grado = storage.GradoMinimo
	}
	if estudiante.Creditos < 0 {
		estudiante.Creditos = 0
	}
	if estudiante.Grado < storage.GradoMinimo {
		estudiante.Grado = storage.GradoMinimo
	}
	if estudiante.Grado > storage.GradoMaximo {
		estudiante.Grado = storage.GradoMaximo
	}

	estudiantes = append(estudiantes, estudiante)
	_ = g.repositorio.Guardar(estudiantes)
	return estudiante
}

func (g *Gestor) BuscarPorID(id int) (models.Estudiante, bool) {
	for _, estudiante := range g.repositorio.Cargar() {
		if estudiante.ID == id {
			return estudiante, true
		}
	}
	return models.Estudiante{}, false
}

func (g *Gestor) Avanzar(id int) (models.Estudiante, bool, error) {
	estudiantes := g.repositorio.Cargar()
	for i, estudiante := range estudiantes {
		if estudiante.ID != id {
			continue
		}
		if !PuedeAvanzar(estudiante) {
			return models.Estudiante{}, false, errors.New("no cumple con los requisitos para avanzar")
		}
		estudiantes[i].Grado++
		estudiantes[i].Creditos = 0
		if err := g.repositorio.Guardar(estudiantes); err != nil {
			return models.Estudiante{}, false, err
		}
		return estudiantes[i], true, nil
	}
	return models.Estudiante{}, false, nil
}

func siguienteID(estudiantes []models.Estudiante) int {
	mayor := 0
	for _, estudiante := range estudiantes {
		if estudiante.ID > mayor {
			mayor = estudiante.ID
		}
	}
	return mayor + 1
}

func PuedeAvanzar(estudiante models.Estudiante) bool {
	return estudiante.Grado >= storage.GradoMinimo && estudiante.Grado < storage.GradoMaximo && estudiante.Creditos >= storage.CreditosRequeridos
}
