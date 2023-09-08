package auxi

import "net/http"

type ServeMux struct {
	*http.ServeMux
}

func NewServeMux() *ServeMux {
	return &ServeMux{
		http.NewServeMux(),
	}
}
