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

			// dto := struct {
			// 	Latitude string `json:"latitude"`
			// }{
			// 	Latitude: c.Cola,
			// }

			helpers.WriteJSON(w, c)
		},
	)
}
