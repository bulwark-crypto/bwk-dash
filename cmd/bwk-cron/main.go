package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Setup database connection.
	db, err := sql.Open("sqlite3", os.ExpandEnv(os.Getenv("DASH_DB")))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Setup to run at interval via cron or use systemd with timer.
	// Get information from node, apis, etc.
	// Store in database.
	// Exit.
}
