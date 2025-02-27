package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"csv_processor/db"
	"csv_processor/server/router"
	v1 "csv_processor/v1"
	"github.com/gorilla/mux"
)

type App struct {
	Router   http.Handler
	Database *sql.DB
}

func New() App {
	app := App{
		Database: db.DB,
	}
	r := router.CSVRouter{Router: mux.NewRouter()}
	v := r.SubPath("/api/v1")
	v1.Route(v)

	cors := os.Getenv("profile") == "prod"
	if !cors {
		r.Use(disableCors)
	}

	app.Router = r
	return app
}

func (a *App) Serve() error {
	port := ":8080"
	srv := &http.Server{
		Handler:      a.Router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("api is available on port %s\n", port)
	return srv.ListenAndServe()
}

func disableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			h.ServeHTTP(w, r)
		},
	)
}
