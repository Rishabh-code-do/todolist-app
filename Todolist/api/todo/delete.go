package todo

import (
	"net/http"

	"github.com/jackc/pgx/v5"
)

func (h *Handler) deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequest(r)
	if err != nil {
		WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.store.DeleteTodo(r.Context(), int32(id))
	if err != nil {
		if err == pgx.ErrNoRows {
			WriteErrorJSON(w, http.StatusNotFound, "Todo not found")
			return
		}
		WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := DeleteResponse{
		ID:      int(res.ID),
		Message: "Todo deleted successfully",
	}

	WriteJSON(w, http.StatusOK, response)
}
