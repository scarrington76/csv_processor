package csv

import (
	"net/http"

	"csv_processor/v1/csv/column"
	"csv_processor/v1/csv/row"
	"github.com/gorilla/mux"
)

func Route(r *mux.Router) {

	r.Path("").
		Methods(http.MethodGet).
		Handler(indexToCtx(index()))

	r.Path("").
		Methods(http.MethodDelete).
		Handler(del())

	r.Path("/row").
		Methods(http.MethodPost).
		Handler(addRow(show()))

	r.Path("").
		Methods(http.MethodPost).
		Handler(create(indexToCtx(index())))

	id := r.PathPrefix("/{id:[0-9]+}").Subrouter()

	rw := id.PathPrefix("/row").Subrouter()
	row.Route(rw)

	c := id.PathPrefix("/column").Subrouter()
	column.Route(c)
}
