package rest

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"token/internal/config"
	"token/internal/transport/rest/routers/iamok"
	"token/internal/transport/rest/routers/token"
	server_http "token/pkg/http"

	"github.com/braintree/manners"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ServerRest struct {
	cfg    *config.Config
	router *chi.Mux
	logger *slog.Logger
}

var server ServerRest

func Init(cfg *config.Config, logger *slog.Logger) *ServerRest {
	server.cfg = cfg
	server.logger = logger
	server.router = server_http.Init(cfg, logger)

	return &server
}

func (s *ServerRest) token() {
	s.router.Route("/token", func(r chi.Router) {
		r.Use(middleware.BasicAuth(
			"ya-token", map[string]string{
				s.cfg.HttpServer.User: s.cfg.HttpServer.Password,
			},
		))
		r.Post("/", token.YaToken(s.logger, &s.cfg.Cloud))
	})
}

func (s *ServerRest) iamok() {
	s.router.Get("/", iamok.IamOK(s.logger))
}

func (s *ServerRest) Run() {

	s.token()
	s.iamok()

	port := fmt.Sprint(":", s.cfg.Port)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go s.listenForShutdown(ch)
	manners.ListenAndServe(port, s.router)
}

func (s *ServerRest) listenForShutdown(ch <-chan os.Signal) {
	<-ch
	fmt.Println("\rshutting down server...")
	manners.Close()
}
