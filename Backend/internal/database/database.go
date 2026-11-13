// A package database connects to postgresql backend
// its  its own package so any part of the backend can import it now
// 
// for example:
// db, err ;= database.Connect()
// instead of: db, err := sql.Open("postgres", dsn) in main.
package database

import (
	"net/url"
	"os"
	"database/sql"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

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
	u := url.URL {
		Scheme: "postgres",
		User:	url.UserPassword(os.Getenv("POSTGRES_USER"),
				os.Getenv("POSTGRES_PASSWORD")),
		Host:	"db:5432",
		Path:	os.Getenv("POSTGRES_DB"),
		RawQuery: "sslmode=disable",
	}
	dsn := u.String()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// OpenGorm is a wrapper to take an existing sql database connection, and reuse it
// there should be a way to init the database with GORM as well, but not sure if its needed here.
func OpenGorm(sqlDB *sql.DB) (*gorm.DB, error) {
	return gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
}