package main

import (
	"fmt"
	"os"

	"github.com/Niflnir/Dreame/internal/database"
	"github.com/pressly/goose"
)

func main() {
	DBCon := database.ConnectToDB()
	migrate_option := os.Args[1]

	if migrate_option == "down" {
		// Rollback migrations to db
		if err := goose.Down(DBCon, "migrations"); err != nil {
			panic(err)
		}
		fmt.Println("Successfully down migrations on database!")
	} else if migrate_option == "up" {
		// Apply migrations to db
		if err := goose.Up(DBCon, "migrations"); err != nil {
			panic(err)
		}
		fmt.Println("Successfully migrated database!")
	}

	DBCon.Close()
}
