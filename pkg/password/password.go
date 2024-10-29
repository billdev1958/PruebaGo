// El paquete `password` proporciona funciones para validar, encriptar y verificar contraseñas.
// Estas funciones incluyen validaciones específicas de requisitos de contraseñas,
// así como el hashing y la verificación utilizando bcrypt.
package password

import (
	"errors"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// ValidatePassword valida si una contraseña cumple con ciertos requisitos de seguridad.
// La contraseña debe:
// - Tener entre 6 y 12 caracteres.
// - Contener al menos una letra mayúscula.
// - Contener al menos una letra minúscula.
// - Contener al menos un número.
// - Contener al menos uno de los siguientes caracteres especiales: @, $, &.
func ValidatePassword(password string) error {

	if len(password) < 6 || len(password) > 12 {
		return errors.New("la contraseña debe tener entre 6 y 12 caracteres")
	}

	if match, _ := regexp.MatchString(`[A-Z]`, password); !match {
		return errors.New("la contraseña debe tener al menos una letra mayúscula")
	}

	if match, _ := regexp.MatchString(`[a-z]`, password); !match {
		return errors.New("la contraseña debe tener al menos una letra minúscula")
	}

	if match, _ := regexp.MatchString(`[0-9]`, password); !match {
		return errors.New("la contraseña debe tener al menos un número")
	}

	if match, _ := regexp.MatchString(`[@$&]`, password); !match {
		return errors.New("la contraseña debe tener al menos uno de los siguientes caracteres: @, $, &")
	}

	return nil
}

// HashPassword genera un hash seguro para la contraseña utilizando bcrypt.
// Devuelve la contraseña hasheada en forma de string, o un error si ocurre algún problema.
func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(bytes), nil
}

// CheckPasswordHash compara una contraseña plana con un hash de contraseña.
// Devuelve `true` si la contraseña coincide con el hash, o `false` si no coincide.
func CheckPasswordHash(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
