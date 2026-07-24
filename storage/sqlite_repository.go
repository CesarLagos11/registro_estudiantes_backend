package storage

import (
	"registro_es/models"

	"gorm.io/gorm"
)

type SqliteEstudianteRepository struct {
	db *gorm.DB
}

func NewSqliteEstudianteRepository(db *gorm.DB) *SqliteEstudianteRepository {
	return &SqliteEstudianteRepository{db: db}
}

func (r *SqliteEstudianteRepository) Cargar() []models.Estudiante {
	var estudiantes []models.Estudiante
	if err := r.db.Order("id").Find(&estudiantes).Error; err != nil {
		return []models.Estudiante{}
	}
	return estudiantes
}

func (r *SqliteEstudianteRepository) Guardar(estudiantes []models.Estudiante) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Exec("DELETE FROM estudiantes").Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(estudiantes) > 0 {
		if err := tx.Create(&estudiantes).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
