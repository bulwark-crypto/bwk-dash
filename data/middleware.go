package data

import (
	"os"

	"github.com/gin-gonic/gin"
)

// Middleware will attach a database connection to the
// the request for use by the handler.
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Setup DB connection.
		db, err := NewSQL(os.ExpandEnv(os.Getenv("DASH_DB")))
		if err != nil {
			c.Error(err)
			return
		}

		// Close database connection.
		defer db.Close()

		c.Set("db", db)
		c.Next()
	}
}
