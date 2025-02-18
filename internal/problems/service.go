package problems

import (
	"net/http"

	"github.com/a-h/templ"
)

type Service struct {
	model ProblemModel
}

func NewService(model ProblemModel) Service {
	return Service{model}
}

func (s Service) Register(mux *http.ServeMux) {
	mux.HandleFunc("/", templ.Handler(List()))
}
