package helper

import (
	"database/sql"
	"log"
)

// CommitOrRollback handles committing or rolling back transactions.
func CommitOrRollback(tx *sql.Tx) {
	if r := recover(); r != nil {
		if err := tx.Rollback(); err != nil {
			log.Printf("Error rolling back transaction: %v\n", err)
			panic(err)
		}
		panic(r) // Re-panic the original error
	} else {
		if err := tx.Commit(); err != nil {
			log.Printf("Error committing transaction: %v\n", err)
			panic(err)
		}
	}
}
