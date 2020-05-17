package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "gqldemopassword"
	DB_NAME     = "gogqldemo"
	SSL_MODE    = "disable"
)

//var once sync.Once

/*
func Connect() (*sql.DB, error) {
	var db *sql.DB
	var err error
	once.Do(func() {
		dbDetails := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", DB_USER, DB_PASSWORD,
			DB_NAME, SSL_MODE)
		db, _ = sql.Open("postgres", dbDetails)
		err = db.Ping()
	})
	return db, err
}
*/

func Connect() (*sql.DB, error) {
	var db *sql.DB
	var err error

	dbDetails := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", DB_USER, DB_PASSWORD,
			DB_NAME, SSL_MODE)
	db, _ = sql.Open("postgres", dbDetails)
	err = db.Ping()

	return db, err
}

func LogAndQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	fmt.Printf("executing %s\n", query)
	return db.Query(query, args...)
}

func MustExec(db *sql.DB, query string, args ...interface{}) {
	_, err := db.Exec(query, args...)
	if err != nil {
		panic(err)
	}
}