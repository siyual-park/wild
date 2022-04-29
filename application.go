package wild

import (
	"net/http"
)

type (
	Application interface {
		Use(middlewareFunc Middleware)
		Start(addr string) error
	}

	app struct {
		h           *Handler
		middlewares []Middleware
	}
)

func New() Application {
	return &app{nil, []Middleware{}}
}

func (a *app) Use(middlewareFunc Middleware) {
	a.middlewares = append(a.middlewares, middlewareFunc)
	h := Compose(a.middlewares)(
		func(c Context) error {
			return nil
		},
	)
	a.h = &h
}

func (a *app) Start(addr string) error {
	return http.ListenAndServe(addr, a)
}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := a.h
	if h != nil {
		_ = (*h)(NewContext(w, r))
	}
}
