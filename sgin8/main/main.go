package main

import (
	"SGin/sgin8"
	"log"
	"net/http"
	"time"
)

func onlyForV2() sgin8.HandlerFunc {
	return func(c *sgin8.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		//c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StateCode, c.Req.RequestURI, time.Since(t))
	}
}
func main() {
	r := sgin8.Default()
	r.GET("/", func(c *sgin8.Context) {
		c.String(http.StatusOK, "Hello wxk\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *sgin8.Context) {
		names := []string{"wxk"}
		c.String(http.StatusOK, names[100]) //访问不到
	})

	r.Run(":9999")
}
