package password_test

import (
	"PruebaGo/pkg/password"
	"testing"
)

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		nombre   string
		password string
		wantErr  bool
	}{
		{"ContraseñaDemasiadoCorta", "A1@", true},
		{"ContraseñaDemasiadoLarga", "A1@aaaaaaaaaaaaa", true},
		{"SinLetraMayúscula", "a1@abcd", true},
		{"SinLetraMinúscula", "A1@12345", true},
		{"SinNúmero", "A@bcdefg", true},
		{"SinCaracterEspecial", "A1bcdefg", true},
		{"ContraseñaVálida", "A1@bcd", false},
	}

	for _, tt := range tests {
		t.Run(tt.nombre, func(t *testing.T) {
			err := password.ValidatePassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePassword() error = %v, quiereError %v", err, tt.wantErr)
			}
		})
	}
}
