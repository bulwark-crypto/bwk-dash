package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/dustinengle/bwk-dash/handler"
	"github.com/dustinengle/bwk-dash/rpc"
	"github.com/dustinengle/bwk-dash/sys"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()
	webdir := os.Getenv("DASH_WEBSITE")

	// Serve static page.
	r.StaticFS("/asset", http.Dir(webdir))
	r.GET("/", func(c *gin.Context) {
		c.File(webdir + "index.html")
	})

	// Setup CORS.
	r.Use(
		cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowCredentials: true,
			AllowHeaders:     []string{"Content-Type"},
			AllowMethods:     []string{"GET", "OPTIONS"},
		}),
	)

	// Setup api routes.
	api := r.Group("/api")
	// Call middleware:
	//  - Attach ENV variables to context
	// 	- Get IP from API
	// 	- Provide RPC Client
	api.Use(
		sys.EnvMiddleware(),
		sys.NetMiddleware(),
		rpc.Middleware(),
	)
	{
		api.GET("/info", handler.GetNodeInfo)
	}

	r.Run(":" + os.Getenv("DASH_PORT"))
}
