package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

func GetConnectionString() string {
	// user -> postgress
	// password -> 12345678
	// host -> localhost
	// port -> 5432
	// db name -> {{dbname}}

	return "user=postgres password=12345678 host=localhost port=5432 dbname=ecommerce"
}

func NewConnection() (*sqlx.DB, error) {
	dbSource := GetConnectionString()
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dbCon, nil
}
