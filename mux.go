package auxi

import "net/http"

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

type MiddlewareChain struct {
	middlewares []Middleware
}

func (mc *MiddlewareChain) AddMiddleware(middleware Middleware) {
	mc.middlewares = append(mc.middlewares, middleware)
}

func (mc *MiddlewareChain) Apply(handler http.HandlerFunc) http.HandlerFunc {
	resultHandler := handler

	for _, middleware := range mc.middlewares {
		resultHandler = middleware(resultHandler)
	}

	return resultHandler
}

func (mc *MiddlewareChain) ApplyToChain(chain MiddlewareChain) *MiddlewareChain {
	return &MiddlewareChain{
		middlewares: append(mc.middlewares, chain.middlewares...),
	}
}

type ServeMux struct {
	*http.ServeMux
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
				http.Error(w, "Method not supported",
					http.StatusMethodNotAllowed)
			}
		}))
}

func NewMiddlewareChain(middlewares ...Middleware) *MiddlewareChain {
	return &MiddlewareChain{
		middlewares: middlewares,
	}
}

func NewServeMux() *ServeMux {
	return &ServeMux{
		http.NewServeMux(),
	}
}

func callMethodHandler(w http.ResponseWriter, r *http.Request, handler func(http.ResponseWriter, *http.Request)) {
	if handler == nil {
		http.NotFound(w, r)
		return
	}

	handler(w, r)
}
