package handler

import (
	"encoding/json"
	"github.com/Collap5e-dev/FlickPick/internal/model"
	"io"
	"net/http"
)

func (h *Handler) MovieAdd(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username := ctx.Value("username").(string)
	NewMovie := model.Movie{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.HandlerError(w, 500, err, "ошибка чтения данных")
		return
	}
	if err := json.Unmarshal(body, &NewMovie); err != nil {
		h.HandlerError(w, 500, err, "ошибка обработки данных")
		return
	}
	defer r.Body.Close()
	err = h.repo.CreateNewMovie(ctx, NewMovie)
	if err != nil {
		h.HandlerError(w, 500, err, "ошибка создания фильма")
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(username))

}
