package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/dustinengle/bwk-dash/data"
	"github.com/dustinengle/bwk-dash/handler"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Setup database connection.
	db, err := data.NewSQL(os.ExpandEnv(os.Getenv("DASH_DB")))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Setup gin framework.
	r := gin.Default()
	webdir := os.ExpandEnv(os.Getenv("DASH_WEBSITE"))

	// Serve static page.
	r.StaticFS("/static", http.Dir(webdir+"/static"))
	r.GET("/", func(c *gin.Context) {
		c.File(webdir + "/index.html")
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
	api.Use(data.Middleware(db))
	{
		api.GET("/info", handler.GetNodeInfo)
	}

	r.Run(":" + os.Getenv("DASH_PORT"))
}
