package row

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

			helpers.WriteJSON(w, c)
		},
	)
}
