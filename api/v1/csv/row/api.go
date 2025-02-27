package row

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"csv_processor/db"
	"csv_processor/models"
	"github.com/gorilla/mux"
)

func del() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			c := ctx.Value("csv").(models.CSV)

			if err := db.DeleteRow(ctx, c.ID); err != nil {
				log.Printf("error deleting csv: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		},
	)
}

func update(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			c := ctx.Value("csv").(models.CSV)
			var dto request

			if c.ID == 0 {
				log.Printf("row does not exist in the db and cannot be updated")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Printf("error reading body: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if err = json.Unmarshal(body, &dto); err != nil {
				log.Printf("error unmarshalling body: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if dto.Name != nil && strings.TrimSpace(*dto.Name) != "" {
				c.Name = *dto.Name
			}

			if dto.Price != nil {
				c.Price = *dto.Price
			}

			if dto.Description != nil && strings.TrimSpace(*dto.Description) != "" {
				c.Description = *dto.Description
			}

			if err = db.UpdateRow(ctx, c); err != nil {
				log.Printf("error applying row update to db: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			ctx = context.WithValue(ctx, "csv", c)

			h.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

func toCtx(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			vars := mux.Vars(r)
			id := vars["id"]
			csv, err := db.GetRow(ctx, id)
			if err != nil && err == sql.ErrNoRows {
				csv = models.CSV{}
			} else if err != nil {
				log.Printf("error retrieving csv: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			ctx = context.WithValue(ctx, "csv", csv)
			h.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}
