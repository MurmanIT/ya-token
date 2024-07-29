package iamok

import (
	"log/slog"
	"net/http"
	"token/internal/transport/rest/response"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func IamOK(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log = log.With(
			slog.String("method", r.Method),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		log.Info("Iam OK")
		render.JSON(w, r, response.ResponseStatus{
			Status: response.StatusOK,
		})
	}
}
