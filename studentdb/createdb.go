package studentdb 

import (
	"database/sql"
)

// createDB - Creates a new PostgreSQL database
func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}
	return nil
}

// ResetDB - Drops the database and creates a new one if one already exists
func ResetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}
	return createDB(db, name)
}
