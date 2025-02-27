package db

import (
	"context"
	"log"

	"csv_processor/models"
)

func PostCSV(ctx context.Context, rows []models.CSV) error {
	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO csv(name, price, description) VALUES($1, $2, $3)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	for _, row := range rows {
		_, err = stmt.Exec(row.Name, row.Price, row.Description)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func GetCSV(ctx context.Context) ([]models.CSV, error) {
	CSVs := []models.CSV{}

	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Error beginning tx")
		return CSVs, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	sqlStatement := `
SELECT * FROM csv`
	rows, err := tx.QueryContext(ctx, sqlStatement)
	if err != nil {
		log.Printf("error getting csv: %v", err)
		return CSVs, err
	}
	defer rows.Close()

	for rows.Next() {
		var csv models.CSV
		if err = rows.Scan(
			&csv.ID, &csv.Name, &csv.Price, &csv.Description,
		); err != nil {
			log.Printf("error scanning")
			return CSVs, err
		}
		CSVs = append(CSVs, csv)
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err = rows.Err(); err != nil {
		log.Printf("error row err")
		return CSVs, err
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		log.Printf("error committing tx")
		return CSVs, err
	}

	return CSVs, nil
}

func DeleteCSV(ctx context.Context) error {
	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Error beginning tx")
		return err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	sqlStatement := `TRUNCATE csv`
	_, err = tx.ExecContext(ctx, sqlStatement)
	if err != nil {
		log.Printf("Error executing update")
		return err
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		log.Printf("error committing tx")
		return err
	}

	return nil
}
