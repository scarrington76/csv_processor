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

func GetRow(ctx context.Context, id string) (models.CSV, error) {
	empty := models.CSV{}
	csv := models.CSV{}

	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Error beginning tx")
		return empty, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	sqlStatement := `
SELECT * FROM csv WHERE id = $1`
	row := tx.QueryRowContext(ctx, sqlStatement, id)
	if err = row.Scan(
		&csv.ID, &csv.Name, &csv.Price, &csv.Description,
	); err != nil {
		log.Printf("error scanning; row may not exist in the db")
		return empty, err
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err = row.Err(); err != nil {
		log.Printf("Error row err")
		return empty, err
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		log.Printf("Error committing tx")
		return empty, err
	}

	return csv, nil
}

func DeleteRow(ctx context.Context, id int64) error {
	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	sqlStatement := `DELETE FROM csv WHERE id=$1`
	_, err = tx.ExecContext(ctx, sqlStatement, id)
	if err != nil {
		log.Printf("error deleting row: %v", err)
		return err
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func UpdateRow(ctx context.Context, c models.CSV) error {
	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	sqlStatement := `UPDATE csv SET (name, description, price) = ($1, $2, $3)
  WHERE id = $4;`
	_, err = tx.ExecContext(ctx, sqlStatement, c.Name, c.Description, c.Price, c.ID)
	if err != nil {
		log.Printf("error updating row: %v", err)
		return err
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
