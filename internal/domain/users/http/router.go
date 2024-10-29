package v1

import (
	"net/http"
)

// UserRoutes define las rutas HTTP para el manejo de usuarios.
// Esta función registra las rutas y las asocia con los métodos correspondientes del handler.
func (h *handler) UserRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/v1/login", h.LoginUser)
	// Registros
	mux.HandleFunc("POST /v1/user", h.RegisterUser)

}
