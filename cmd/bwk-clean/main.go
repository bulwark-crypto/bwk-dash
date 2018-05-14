package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Remove database file.
	err := os.Remove(os.ExpandEnv(os.Getenv("DASH_DB")))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database removed!")
}
