package router

import (
	"github.com/gorilla/mux"
)

type CSVRouter struct {
	*mux.Router
}

func (r CSVRouter) SubPath(s string) CSVRouter {
	return CSVRouter{r.PathPrefix(s).Subrouter()}
}
