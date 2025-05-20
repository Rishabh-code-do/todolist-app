package todo

import (
	"net/http"
)

func (h *Handler) markTodoCompleted(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequest(r)
	if err != nil {
		WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.store.MarkTodoAsCompleted(r.Context(), int32(id))

	if err != nil {
		WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	response := ResponseTodo{
		ID:          int(res.ID),
		Title:       res.Title,
		Description: res.Description.String,
		Status:      string(res.Status.TaskStatus),
		DueDate:     res.Duedate.Time,
		CreatedAt:   res.Createdat.Time,
		UpdatedAt:   res.Updatedat.Time,
	}
	WriteJSON(w, http.StatusOK, response)
}
