package problems

import (
	"github.com/donseba/go-htmx"
	"net/http"
)

type Service struct {
	model ProblemModel
	htmx     *htmx.HTMX
}

func NewService(model ProblemModel) Service {
	return Service{model, htmx.New()}
}

func (s Service) Register(mux *http.ServeMux) {
	mux.HandleFunc("/", s.listPage)
}
