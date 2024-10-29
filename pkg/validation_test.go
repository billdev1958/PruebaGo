package validation_test

import (
	validation "PruebaGo/pkg"
	"testing"
)

func TestValidateUser(t *testing.T) {
	tests := []struct {
		name     string
		username string
		phone    string
		email    string
		pass     string
		wantErr  bool
	}{
		{"NombreUsuarioVacio", "", "1234567890", "test@example.com", "A1@bcd", true},
		{"TelefonoMenosDe10Digitos", "john_doe", "123456789", "test@example.com", "A1@bcd", true},
		{"TelefonoMasDe10Digitos", "john_doe", "12345678901", "test@example.com", "A1@bcd", true},
		{"TelefonoConCaracteresInvalidos", "john_doe", "12345abcd0", "test@example.com", "A1@bcd", true},
		{"EmailInvalido", "john_doe", "1234567890", "invalid-email", "A1@bcd", true},
		{"ContraseñaInvalida", "john_doe", "1234567890", "test@example.com", "abc", true},
		{"DatosValidos", "john_doe", "1234567890", "test@example.com", "A1@bcd", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.ValidateUser(tt.username, tt.phone, tt.email, tt.pass)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateLoginData(t *testing.T) {
	tests := []struct {
		name       string
		identifier string
		pass       string
		wantErr    bool
	}{
		{"IdentificadorVacio", "", "A1@bcd", true},
		{"ContraseñaVacia", "john_doe", "", true},
		{"DatosValidos", "john_doe", "A1@bcd", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.ValidateLoginData(tt.identifier, tt.pass)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateLoginData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
