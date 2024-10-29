package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var secretKey []byte

// funcion para obtener la clave secreta del token de las variables de entorno
func init() {
	secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(secretKey) == 0 {
		panic("JWT_SECRET_KEY is not set in environment variables")
	}
}

// Estructura para definir los claims del token
type Claims struct {
	ID       uuid.UUID
	Username string
	jwt.RegisteredClaims
}

// Genera un token almacenando el ID y el username dentro de el
func GenerateJWT(id uuid.UUID, username string) (string, error) {
	claims := Claims{
		ID:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
			Issuer:    "prueba_go",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	fmt.Println("Token being generated:", tokenString)
	return tokenString, nil
}

// ValidateJWT valida un token JWT y devuelve las reclamaciones (claims) si el token es válido.
// Esta función utiliza el método de firma HMAC para garantizar la seguridad de los tokens.
// Recibe el token JWT como string, y devuelve un puntero a `Claims` y un posible error.
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Validar que el método de firma sea HMAC.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		fmt.Println("Error parsing token:", err)
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		fmt.Println("Token is invalid:", tokenString)
		return nil, fmt.Errorf("invalid token: %w", jwt.ErrSignatureInvalid)
	}

	fmt.Println("Token is valid, claims:", claims)
	return claims, nil
}
