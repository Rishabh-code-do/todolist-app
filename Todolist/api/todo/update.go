package todo

import (
	"net/http"
	db "todolist/db/sqlc"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Handler) updateTodo(w http.ResponseWriter, r *http.Request) {

	id, err := ParseIDFromRequest(r)
	if err != nil {
		WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	var req UpdateTodoRequest
	if err := ReadJSON(w, r, &req); err != nil {
		return
	}

	if !ValidateStruct(w, req) {
		return
	}

	res, err := h.store.UpdateTodo(r.Context(), db.UpdateTodoParams{
		Title:       req.Title,
		Description: pgtype.Text{String: req.Description, Valid: true},
		Duedate:     pgtype.Date{Time: req.DueDate, Valid: true},
		ID:          int32(id),
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			WriteErrorJSON(w, http.StatusNotFound, "Todo not found")
			return
		}
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
