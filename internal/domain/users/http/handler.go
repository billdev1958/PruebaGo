package v1

import (
	user "PruebaGo/internal/domain/users"
	"PruebaGo/internal/domain/users/http/models"
	"encoding/json"
	"log"
	"net/http"
)

type handler struct {
	uc user.Usecase
}

func NewHandler(uc user.Usecase) *handler {
	return &handler{uc: uc}
}

func (h *handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var request models.LoginRequest

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("Error al decodificar la solicitud: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Llamada al caso de uso
	loginResponse, err := h.uc.Login(r.Context(), request)
	if err != nil {
		log.Printf("Error en el proceso de login: %v", err)
		response := models.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := models.Response{
		Status:  "success",
		Message: "User logged in successfully",
		Data:    loginResponse,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error al codificar la respuesta: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

}

func (h *handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request models.RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	log.Printf("Decoded request: %+v", request)

	user, err := h.uc.RegisterUser(r.Context(), request)
	if err != nil {
		response := models.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := models.Response{
		Status:  "success",
		Message: "User registered successfully",
		Data:    user,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
