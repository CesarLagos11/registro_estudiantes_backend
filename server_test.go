package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"registro_es/models"
	"registro_es/storage"

	"github.com/labstack/echo/v4"
)

func TestCrearEstudianteAPI(t *testing.T) {
	tempDir := t.TempDir()
	t.Chdir(tempDir)

	e := echo.New()
	payload := `{"nombre":"Ana","grado":4,"creditos":7}`
	req := httptest.NewRequest(http.MethodPost, "/estudiantes", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := CrearEstudianteAPI(c); err != nil {
		t.Fatalf("crear estudiante: %v", err)
	}

	if rec.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusCreated)
	}

	contenido, err := os.ReadFile(filepath.Join(tempDir, storage.ArchivoEstudiantes))
	if err != nil {
		t.Fatalf("leer archivo: %v", err)
	}

	var estudiantes []models.Estudiante
	if err := json.Unmarshal(contenido, &estudiantes); err != nil {
		t.Fatalf("decodificar JSON: %v", err)
	}
	if len(estudiantes) != 1 || estudiantes[0].Nombre != "Ana" {
		t.Fatalf("el estudiante no se persistió correctamente: %s", contenido)
	}
}

func TestRutaRaizDevuelveInfo(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := RootHandler(c); err != nil {
		t.Fatalf("root handler: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if !strings.Contains(rec.Body.String(), "health") {
		t.Fatalf("body = %q, want endpoint info", rec.Body.String())
	}
}
