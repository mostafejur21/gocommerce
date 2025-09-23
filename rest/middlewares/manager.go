package middlewares

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddleware []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddleware: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddleware = append(mngr.globalMiddleware, middlewares...)
}

func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler

	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {
	h := handler
	for _, middleware := range mngr.globalMiddleware {
		h = middleware(h)
	}
	return h
}
