package db

import (
	"context"
	"log"

	"csv_processor/models"
)

func AddRow(ctx context.Context, c models.CSV) (models.CSV, error) {
	empty := models.CSV{}
	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		return empty, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	var id int64
	sqlStatement := `
INSERT INTO csv (name, price, description) VALUES ($1, $2, $3) RETURNING id`
	if err = tx.QueryRowContext(
		ctx, sqlStatement, c.Name, c.Price, c.Description,
	).Scan(&id); err != nil {
		log.Printf("error inserting row: %v", err)
		return empty, err
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return empty, err
	}

	c.ID = id

	return c, nil
}
