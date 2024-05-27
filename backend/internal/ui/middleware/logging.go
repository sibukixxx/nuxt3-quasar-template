package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger is a simple middleware to log requests.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print("Latency: ", latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println("Status: ", status)
	}
}
