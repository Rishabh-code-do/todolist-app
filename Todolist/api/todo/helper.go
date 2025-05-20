package todo

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ParseIDFromRequest(r *http.Request) (int32, error) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		return 0, errors.New("missing ID in URL")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return 0, errors.New("invalid ID: must be a positive integer")
	}

	return int32(id), nil
}

func ReadJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(dst)
	if err != nil {
		WriteErrorJSON(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return err
	}
	return nil
}

func ValidateStruct(w http.ResponseWriter, data any) bool {
	if err := validate.Struct(data); err != nil {
		WriteErrorJSON(w, http.StatusBadRequest, "validation failed: "+err.Error())
		return false
	}
	return true
}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func WriteErrorJSON(w http.ResponseWriter, status int, message string) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	WriteJSON(w, status, errorResponse{Error: message})
}
