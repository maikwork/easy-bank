package server

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/maikwork/balanceUserAvito/internall/handler"
	"github.com/maikwork/balanceUserAvito/internall/model"
	"github.com/maikwork/balanceUserAvito/internall/repository"
)

type Server struct {
	Cnf       model.Config
	UserStore repository.UserStore
}

func New(store repository.UserStore) *Server {
	return &Server{
		UserStore: store,
	}
}

func (s *Server) Run() {
	mux := http.NewServeMux()

	h := handler.NewHandeler(s.UserStore)

	mux.HandleFunc("/", h.Handler)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.WithError(err).Fatal("server don't run")
	}
}
