package sys

import (
	"os"

	"github.com/gin-gonic/gin"
)

// EnvMiddleware will load the donation address
// and other ENV values into context if needed.
func EnvMiddleware() gin.HandlerFunc {
	donation := os.Getenv("DASH_DONATION_ADDRESS")

	return func(c *gin.Context) {
		c.Set("donation", donation)
		c.Next()
	}
}

// NetMiddleware will fetch the ip address from
// an api and will provide it in the request
// context.  Each time the api starts it should
// perform this action only once per client.
func NetMiddleware() gin.HandlerFunc {
	ip, err := GetIP()
	if err != nil {
		ip = "Unknown"
	}

	return func(c *gin.Context) {
		c.Set("ip", ip)
		c.Next()
	}
}
