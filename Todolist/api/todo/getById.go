package todo

import (
	"net/http"

	"github.com/jackc/pgx/v5"
)

func (h *Handler) getTodosById(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequest(r)
	if err != nil {
		WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	todo, err := h.store.GetTodoById(r.Context(), int32(id))
	if err != nil {
		if err == pgx.ErrNoRows {
			WriteErrorJSON(w, http.StatusNotFound, "Todo not found")
			return
		}
		WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	response := ResponseTodo{
		ID:          int(todo.ID),
		Title:       todo.Title,
		Description: todo.Description.String,
		Status:      string(todo.Status.TaskStatus),
		CreatedAt:   todo.Createdat.Time,
		UpdatedAt:   todo.Updatedat.Time,
		DueDate:     todo.Duedate.Time,
	}

	WriteJSON(w, http.StatusOK, response)
}
