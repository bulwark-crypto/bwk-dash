package handler

import (
	"github.com/gin-gonic/gin"
)

// GetNodeInfo will return all the information
// needed to be displayed on the dashboard.
func GetNodeInfo(c *gin.Context) {
	// Get DB connection.
	// Pull last entry and return.

	// Return as json to client.
	c.JSON(200, "")
}
