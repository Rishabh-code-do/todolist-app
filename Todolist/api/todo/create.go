package todo

import (
	"net/http"
	db "todolist/db/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Handler) createTodo(w http.ResponseWriter, r *http.Request) {

	var req CreateTodoRequest
	if err := ReadJSON(w, r, &req); err != nil {
		return
	}

	if !ValidateStruct(w, req) {
		return
	}

	res, err := h.store.CreateTodo(r.Context(), db.CreateTodoParams{
		Title:       req.Title,
		Description: pgtype.Text{String: req.Description, Valid: true},
		Status:      db.NullTaskStatus{TaskStatus: db.TaskStatusPENDING, Valid: true},
		Duedate:     pgtype.Date{Time: req.DueDate, Valid: true},
	})
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

	WriteJSON(w, http.StatusCreated, response)
}
