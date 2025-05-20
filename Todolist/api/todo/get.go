package todo

import (
	"net/http"
)

func (h *Handler) getTodos(w http.ResponseWriter, r *http.Request) {

	todolist, err := h.store.GetAllTodo(r.Context())
	if err != nil {
		WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	WriteJSON(w, http.StatusOK, todolist)
}
