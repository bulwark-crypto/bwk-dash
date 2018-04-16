package rpc

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Middleware will attach a reference to the
// rpc client to the request context.
func Middleware() gin.HandlerFunc {
	// Setup rpc connection.
	node, err := NewRPC()
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		// Attach rpc client to context.
		c.Set("rpc", node)

		// Call next in chain.
		c.Next()
	}
}
