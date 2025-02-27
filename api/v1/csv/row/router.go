package row

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Route(r *mux.Router) {
	r.Use(toCtx)

	r.Path("").
		Methods(http.MethodDelete).
		Handler(del())

	r.Path("").
		Methods(http.MethodPut).
		Handler(update(show()))

	r.Path("").
		Methods(http.MethodGet).
		Handler(show())
}
