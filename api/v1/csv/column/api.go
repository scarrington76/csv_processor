package column

import (
	"context"
	"net/http"

	"csv_processor/models"
)

func toCtx(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			csv := models.CSV{
				ID:          1,
				Name:        "something",
				Price:       0.99,
				Description: "description",
			}

			// csv, err := db.GetCSV(ctx)
			// if err != nil {
			// 	fmt.Printf("error retrieving csv: %v", err)
			// 	return
			// }

			// csv.Cola = "something"

			ctx = context.WithValue(ctx, "csv", csv)

			h.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

func del(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			csv := models.CSV{
				ID:          1,
				Name:        "something",
				Price:       0.99,
				Description: "description",
			}

			// csv, err := db.GetCSV(ctx)
			// if err != nil {
			// 	fmt.Printf("error retrieving csv: %v", err)
			// 	return
			// }

			// csv.Cola = "something"

			ctx = context.WithValue(ctx, "csv", csv)

			h.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

func update(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			csv := models.CSV{
				ID:          1,
				Name:        "something",
				Price:       0.99,
				Description: "description",
			}

			// csv, err := db.GetCSV(ctx)
			// if err != nil {
			// 	fmt.Printf("error retrieving csv: %v", err)
			// 	return
			// }

			// csv.Cola = "something"

			ctx = context.WithValue(ctx, "csv", csv)

			h.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}
