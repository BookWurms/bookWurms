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

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
func Connect() (*gorm.DB, error) {
	u := url.URL{
		Scheme: "postgres",
		User: url.UserPassword(os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD")),
		Host:     "db:5432",
		Path:     os.Getenv("POSTGRES_DB"),
		RawQuery: "sslmode=disable",
	}
	dsn := u.String()

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gormDB, nil
}
