package token

import (
	"log/slog"
	"net/http"
	"token/internal/config"
	"token/internal/transport/rest/response"

	clientclink "token/pkg/client_clink"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type ResponseToken struct {
	response.ResponseStatus
	Token string `json:"token"`
}

func YaToken(log *slog.Logger, cloud *config.Cloud) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log = log.With(
			slog.String("method", r.Method),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		client := clientclink.ClientC{}
		token, err := client.Clink(cloud.UrlCloud, cloud.Token)
		if err != nil {
			log.Error("Error getting token", slog.Any("error", err))
			render.JSON(w, r, response.Error("Error getting token"))
			return
		}

		render.JSON(w, r, ResponseToken{
			ResponseStatus: response.Ok(),
			Token:          token.BasicToken,
		})
	}
}
