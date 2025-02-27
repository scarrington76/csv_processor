package column

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Route(r *mux.Router) {

	r.Path("").
		Methods(http.MethodDelete).
		Handler(del(show()))

	r.Path("").
		Methods(http.MethodPut).
		Handler(update(show()))

	r.Path("").
		Methods(http.MethodGet).
		Handler(show())
}
