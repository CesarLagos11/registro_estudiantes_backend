package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"registro_es/estudiantes"
	"registro_es/handlers"
	"registro_es/models"
	"registro_es/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "api" {
		e := echo.New()
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:8080", "http://localhost:8081", "http://127.0.0.1:8080", "http://127.0.0.1:8081"},
			AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}))
		ConfigurarRutas(e)
		fmt.Println("Servidor escuchando en http://localhost:1323")
		if err := e.Start(":1323"); err != nil {
			fmt.Println(err)
		}
		return
	}

	reader := bufio.NewReader(os.Stdin)
	handlers.EjecutarMenu(reader)
}

func ConfigurarRutas(e *echo.Echo) {
	e.GET("/", RootHandler)
	e.GET("/health", HealthCheck)
	e.GET("/estudiantes", ListarEstudiantesAPI)
	e.POST("/estudiantes", CrearEstudianteAPI)
	e.GET("/estudiantes/:id", ObtenerEstudianteAPI)
	e.POST("/estudiantes/:id/avanzar", AvanzarEstudianteAPI)
}

func RootHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "API de registro de estudiantes",
		"endpoints": []string{"/health", "/estudiantes", "/estudiantes/:id", "/estudiantes/:id/avanzar"},
	})
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func ListarEstudiantesAPI(c echo.Context) error {
	repo := storage.NewArchivoEstudianteRepository(storage.ArchivoEstudiantes)
	gestor := estudiantes.NewGestor(repo)
	return c.JSON(http.StatusOK, gestor.Listar())
}

func CrearEstudianteAPI(c echo.Context) error {
	var req struct {
		Nombre    string  `json:"nombre"`
		Edad      *int    `json:"edad,omitempty"`
		Correo    *string `json:"correo,omitempty"`
		NumCuenta *string `json:"numCuenta,omitempty"`
		Grado     int     `json:"grado"`
		Creditos  int     `json:"creditos"`
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "datos inválidos")
	}

	nombre := strings.TrimSpace(req.Nombre)
	if nombre == "" {
		nombre = "No registrado"
	}
	if req.Grado == 0 {
		req.Grado = storage.GradoMinimo
	}
	if req.Grado < storage.GradoMinimo || req.Grado > storage.GradoMaximo {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("grado inválido: debe estar entre %d y %d", storage.GradoMinimo, storage.GradoMaximo))
	}
	if req.Creditos < 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "los créditos no pueden ser negativos")
	}

	repo := storage.NewArchivoEstudianteRepository(storage.ArchivoEstudiantes)
	gestor := estudiantes.NewGestor(repo)
	estudiante := gestor.Registrar(models.Estudiante{
		Nombre:    nombre,
		Edad:      req.Edad,
		Correo:    req.Correo,
		NumCuenta: req.NumCuenta,
		Grado:     req.Grado,
		Creditos:  req.Creditos,
	})

	return c.JSON(http.StatusCreated, estudiante)
}

func ObtenerEstudianteAPI(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id inválido")
	}

	repo := storage.NewArchivoEstudianteRepository(storage.ArchivoEstudiantes)
	gestor := estudiantes.NewGestor(repo)
	if estudiante, ok := gestor.BuscarPorID(id); ok {
		return c.JSON(http.StatusOK, estudiante)
	}

	return echo.NewHTTPError(http.StatusNotFound, "estudiante no encontrado")
}

func AvanzarEstudianteAPI(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id inválido")
	}

	repo := storage.NewArchivoEstudianteRepository(storage.ArchivoEstudiantes)
	gestor := estudiantes.NewGestor(repo)
	if estudiante, ok, err := gestor.Avanzar(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if ok {
		return c.JSON(http.StatusOK, estudiante)
	}

	return echo.NewHTTPError(http.StatusNotFound, "estudiante no encontrado")
}
