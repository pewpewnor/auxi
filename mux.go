package auxi

import (
	"net/http"

	"github.com/pewpewnor/auxi/respond"
)

type Middleware func(next http.HandlerFunc) http.HandlerFunc

type MethodHandlers struct {
	GET     func(http.ResponseWriter, *http.Request)
	POST    func(http.ResponseWriter, *http.Request)
	PUT     func(http.ResponseWriter, *http.Request)
	DELETE  func(http.ResponseWriter, *http.Request)
	PATCH   func(http.ResponseWriter, *http.Request)
	HEAD    func(http.ResponseWriter, *http.Request)
	OPTIONS func(http.ResponseWriter, *http.Request)
}

type ServeMux struct {
	*http.ServeMux
	middlewareChain *MiddlewareChain
}

func NewServeMux() *ServeMux {
	return &ServeMux{
		http.NewServeMux(),
		NewMiddlewareChain(),
	}
}

func (mux *ServeMux) Use(middleware Middleware) {
	mux.middlewareChain.AddMiddleware(middleware)
}

func (mux *ServeMux) UseChain(middlewareChain *MiddlewareChain) {
	mux.middlewareChain.ApplyToChain(middlewareChain)
}

func (mux *ServeMux) HandleMethods(pattern string, methodHandlers MethodHandlers) {
	mux.HandleFunc(pattern, http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				callMethodHandler(w, r, methodHandlers.GET)
			case http.MethodPost:
				callMethodHandler(w, r, methodHandlers.POST)
			case http.MethodPut:
				callMethodHandler(w, r, methodHandlers.PUT)
			case http.MethodDelete:
				callMethodHandler(w, r, methodHandlers.DELETE)
			case http.MethodPatch:
				callMethodHandler(w, r, methodHandlers.PATCH)
			case http.MethodHead:
				callMethodHandler(w, r, methodHandlers.HEAD)
			case http.MethodOptions:
				callMethodHandler(w, r, methodHandlers.OPTIONS)
			default:
				respond.JSON(w, respond.SError("Method not supported"),
					http.StatusMethodNotAllowed)
			}
		}))
}

func (mux *ServeMux) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	mux.ServeMux.HandleFunc(pattern, mux.middlewareChain.Apply(handler))
}

type MiddlewareChain struct {
	middlewares []Middleware
}

func NewMiddlewareChain(middlewares ...Middleware) *MiddlewareChain {
	return &MiddlewareChain{
		middlewares: middlewares,
	}
}

func (c *MiddlewareChain) AddMiddleware(middleware Middleware) {
	c.middlewares = append(c.middlewares, middleware)
}

func (c *MiddlewareChain) Apply(handler http.HandlerFunc) http.HandlerFunc {
	resultHandler := handler

	for _, middleware := range c.middlewares {
		resultHandler = middleware(resultHandler)
	}

	return resultHandler
}

func (c *MiddlewareChain) ApplyToChain(chain *MiddlewareChain) *MiddlewareChain {
	return &MiddlewareChain{
		middlewares: append(c.middlewares, chain.middlewares...),
	}
}

func callMethodHandler(w http.ResponseWriter, r *http.Request, handler func(http.ResponseWriter, *http.Request)) {
	if handler == nil {
		http.NotFound(w, r)
		return
	}

	http.HandlerFunc(handler).ServeHTTP(w, r)
}
