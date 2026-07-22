package handlers

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"registro_es/storage"
)

var (
	regexCorreo = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	regexCuenta = regexp.MustCompile(`^\d{2}-\d{4}-\d{5}$`)
)

func LeerEdad(reader *bufio.Reader, mensaje string) *int {
	for {
		texto := LeerLinea(reader, mensaje)
		if texto == "" {
			return nil
		}
		if !EsEdadValida(texto) {
			fmt.Println("Edad inválida. Ingrese solo números.")
			continue
		}

		valor, _ := strconv.Atoi(texto)
		return &valor
	}
}

func EsEdadValida(texto string) bool {
	if texto == "" {
		return false
	}

	_, err := strconv.Atoi(texto)
	return err == nil
}

func LeerCorreo(reader *bufio.Reader, mensaje string) *string {
	for {
		texto := LeerLinea(reader, mensaje)
		if texto == "" {
			return nil
		}
		if !EsCorreoValido(texto) {
			fmt.Println("Correo inválido. Ejemplo: nombre@dominio.com")
			continue
		}
		return &texto
	}
}

func EsCorreoValido(texto string) bool {
	return regexCorreo.MatchString(texto)
}

func LeerNumCuenta(reader *bufio.Reader, mensaje string) *string {
	for {
		texto := LeerLinea(reader, mensaje)
		if texto == "" {
			return nil
		}
		if !EsNumCuentaValido(texto) {
			fmt.Println("Número de cuenta inválido. Debe tener el formato 12-1999-12345.")
			continue
		}
		return &texto
	}
}

func EsNumCuentaValido(texto string) bool {
	return regexCuenta.MatchString(texto)
}

func LeerGrado(reader *bufio.Reader, mensaje string) int {
	for {
		texto := LeerLinea(reader, mensaje)
		if texto == "" {
			return storage.GradoMinimo
		}
		valor, err := strconv.Atoi(texto)
		if err != nil || valor < storage.GradoMinimo || valor > storage.GradoMaximo {
			fmt.Println("Grado inválido. Ingrese un número entre 1 y 9.")
			continue
		}
		return valor
	}
}

func LeerCreditos(reader *bufio.Reader, mensaje string) int {
	for {
		texto := LeerLinea(reader, mensaje)
		if texto == "" {
			return 0
		}
		valor, err := strconv.Atoi(texto)
		if err != nil || valor < 0 {
			fmt.Println("Créditos inválidos. Ingrese un número mayor o igual a 0.")
			continue
		}
		return valor
	}
}

func LeerLinea(reader *bufio.Reader, mensaje string) string {
	fmt.Print(mensaje)
	texto, _ := reader.ReadString('\n')
	return strings.TrimSpace(texto)
}

func LeerOpcion(reader *bufio.Reader, mensaje string) int {
	for {
		texto := LeerLinea(reader, mensaje)
		if texto == "" {
			return 0
		}

		valor, err := strconv.Atoi(texto)
		if err == nil {
			return valor
		}

		fmt.Println("Opción inválida. Intente nuevamente.")
	}
}
