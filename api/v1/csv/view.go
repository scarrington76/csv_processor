package csv

import (
	"net/http"

	"csv_processor/helpers"
	"csv_processor/models"
)

func show() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			c := ctx.Value("csv").(models.CSV)

			dto := models.CSV{
				ID:          c.ID,
				Name:        c.Name,
				Description: c.Description,
				Price:       c.Price,
			}

			helpers.WriteJSON(w, dto)
		},
	)
}

func index() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			rows := ctx.Value("csv").([]models.CSV)

			dto := []any{}

			for _, rw := range rows {
				l := models.CSV{
					ID:          rw.ID,
					Name:        rw.Name,
					Description: rw.Description,
					Price:       rw.Price,
				}
				dto = append(dto, l)
			}

			helpers.WriteJSON(w, dto)
		},
	)
}
