package csv

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"

	"csv_processor/db"
	"csv_processor/helpers"
	"csv_processor/models"
)

func indexToCtx(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			csv, err := db.GetCSV(ctx)
			if err != nil && errors.Is(err, sql.ErrNoRows) {
				csv = []models.CSV{}
			} else if err != nil {
				log.Printf("error retireving csv: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			ctx = context.WithValue(ctx, "csv", csv)
			h.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

func create(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			body, err := io.ReadAll(r.Body)
			var dto requestCSV

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

			if dto.Filename == nil || strings.TrimSpace(*dto.Filename) == "" {
				log.Print("invalid filename provided")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			csv, err := helpers.ImportCSVtoStruct(*dto.Filename)
			if err != nil {
				log.Printf("error importing csv: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if err = db.PostCSV(ctx, csv); err != nil {
				log.Printf("error applying csv to db: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			h.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

func del() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			if err := db.DeleteCSV(ctx); err != nil {
				log.Printf("error applying csv to db: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		},
	)
}

func addRow(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			body, err := io.ReadAll(r.Body)
			var dto requestRow

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

			if dto.Name == nil || strings.TrimSpace(*dto.Name) == "" {
				log.Print("invalid name provided")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if dto.Price == nil {
				log.Print("invalid filename provided")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if dto.Description == nil || strings.TrimSpace(*dto.Description) == "" {
				log.Print("invalid description provided")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			c := models.CSV{
				Name:        *dto.Name,
				Price:       *dto.Price,
				Description: *dto.Description,
			}

			csv, err := db.AddRow(ctx, c)
			if err != nil {
				log.Printf("error applying row to db: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			ctx = context.WithValue(ctx, "csv", csv)

			h.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}
