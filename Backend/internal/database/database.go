// A package database connects to postgresql backend
// its  its own package so any part of the backend can import it now
// 
// for example:
// db, err ;= database.Connect()
// instead of: db, err := sql.Open("postgres", dsn) in main.
package database

import (
	"fmt"
	"os"
	"database/sql"

	_ "github.com/lib/pq"
)

/* 
Why return nil, err instead of fatal.log?
because if we want to run a lot of tests later.
fatal.log will end the program early, wheras returning only the error allows
for the tests to continue, also log.fatal skips any cleanup like 
defer statements or db.close()
Kind of like a try catch, so the function return the errors, 
and the main decides if its worth exiting.
*/
func Connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@db:5432/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}