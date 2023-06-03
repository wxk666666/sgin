package middleware

import (
	"SGin/sgin8"
	"log"
	"time"
)

func Logger() sgin8.HandlerFunc {
	return func(c *sgin8.Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StateCode, c.Req.RequestURI, time.Since(t))
	}
}
