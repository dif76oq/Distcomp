package handler

import (
	"github.com/Khmelov/Distcomp/251003/truhan/labX/publisher/internal/handler/http"
	"github.com/Khmelov/Distcomp/251003/truhan/labX/publisher/internal/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	HTTP *mux.Router
}

func New(srv service.Service) Handler {
	return Handler{
		HTTP: http.New(srv),
	}
}
