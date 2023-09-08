package auxi

import "net/http"

type MethodHandlers struct {
	GET     func(http.ResponseWriter, *http.Request)
	POST    func(http.ResponseWriter, *http.Request)
	PUT     func(http.ResponseWriter, *http.Request)
	DELETE  func(http.ResponseWriter, *http.Request)
	PATCH   func(http.ResponseWriter, *http.Request)
	HEAD    func(http.ResponseWriter, *http.Request)
	OPTIONS func(http.ResponseWriter, *http.Request)
}

func callMethodHandler(w http.ResponseWriter, r *http.Request, handler func(http.ResponseWriter, *http.Request)) {
	if handler == nil {
		http.NotFound(w, r)
		return
	}

	handler(w, r)
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
