package models

type Estudiante struct {
	ID        int     `json:"id"`
	Nombre    string  `json:"nombre"`
	Edad      *int    `json:"edad,omitempty"`
	Correo    *string `json:"correo,omitempty"`
	NumCuenta *string `json:"numCuenta,omitempty"`
	Grado     int     `json:"grado"`
	Creditos  int     `json:"creditos"`
}
