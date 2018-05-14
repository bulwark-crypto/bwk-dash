package data

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

// Middleware will attach a database connection to
// the request for use in the handler.
func Middleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add to context for use in handler.
		c.Set("db", db)

		// Go to next middleware or route.
		c.Next()
	}
}
