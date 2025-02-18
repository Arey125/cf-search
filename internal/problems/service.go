package problems

import "net/http"

type Service struct {
	model ProblemModel
}

func NewService(model ProblemModel) Service {
	return Service{model}
}

func (s Service) Register(mux *http.ServeMux) {
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Test"))
    })
}
