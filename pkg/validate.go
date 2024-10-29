// El paquete validation contiene algunas utilidades para validar la información del usuario.
// Incluye validaciones para el registro y el inicio de sesión de usuarios.
package validation

import (
	"PruebaGo/pkg/password"
	"errors"
	"fmt"
	"net/mail"
	"regexp"
)

func ValidateUser(username, phone, email, pass string) error {
	// Validar que el nombre de usuario no esté vacío
	if username == "" {
		return errors.New("username no debe estar vacío")
	}

	// Validar que el número de teléfono tenga exactamente 10 dígitos
	if len(phone) != 10 {
		return errors.New("el número de teléfono debe tener exactamente 10 dígitos")
	}

	// Validar que el número de teléfono contenga solo dígitos
	if match, _ := regexp.MatchString(`^\d{10}$`, phone); !match {
		return errors.New("el número de teléfono debe contener solo dígitos")
	}

	// Validar que el email tenga un formato válido
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("email debe tener un formato válido")
	}

	// Validar la contraseña utilizando la función ValidatePassword
	if err := password.ValidatePassword(pass); err != nil {
		return fmt.Errorf("contraseña inválida: %w", err)
	}

	return nil
}

// ValidateLoginData valida los datos proporcionados durante el inicio de sesión.
// Comprueba que tanto el identificador (username o email) como la contraseña estén presentes.
func ValidateLoginData(identifier, pass string) error {
	if identifier == "" {
		return errors.New("debe proporcionar un identificador (username o email)")
	}
	if pass == "" {
		return errors.New("falta campo contraseña")
	}

	return nil
}
