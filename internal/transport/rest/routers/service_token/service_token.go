package servicetoken

import (
	"log/slog"
	"net/http"
	"token/internal/config"
	"token/internal/transport/rest/response"
	clientclink "token/pkg/client_clink"
	jwttoken "token/pkg/jwt_token"

	"github.com/go-chi/render"
)

type ResponseServiceToken struct {
	response.ResponseStatus
	Token string `json:"token"`
}

func ServiceToken(log *slog.Logger, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ch := make(chan string)

		go getToken(log, cfg, ch)

		token := <-ch

		render.JSON(w, r, ResponseServiceToken{
			ResponseStatus: response.Ok(),
			Token:          token,
		})
	}
}

func getToken(log *slog.Logger, cfg *config.Config, ch chan string) {
	configJWT := jwttoken.ConfigJWTToken{
		ServiceAccount: &cfg.ServiceAccount,
		Logger:         log,
		Url:            cfg.Cloud.UrlCloud,
	}
	token, err := jwttoken.Generate(&configJWT)
	if err != nil {
		log.Error("Error generating token", slog.Any("error", err))
		return
	}

	client := clientclink.ClientC{}
	yaToken, errToken := client.ClinkService(cfg.Cloud.UrlCloud, token)

	if errToken != nil {
		log.Error("Error getting token", slog.Any("error", errToken))
		return
	}

	ch <- yaToken.BasicToken
}
