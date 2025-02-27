package v1

import (
	"csv_processor/server/router"
	"csv_processor/v1/csv"
)

func Route(r router.CSVRouter) {

	c := r.PathPrefix("/csv").Subrouter()
	csv.Route(c)
}
