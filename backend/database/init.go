package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/pressly/goose"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "dreame"
)

var DBCon *sql.DB

func ConnectToDB() *sql.DB {
  var err error
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DBCon, err = sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

  // Test DB connection
  err = DBCon.Ping()
  if err != nil {
    log.Fatal(err)
  }
	fmt.Printf("Successfully connected to database:%s!\n", dbname)

	// Set SQL dialect to postgresql
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	fmt.Println("Successfully changed dialect to postgres!")

  return DBCon
}
