package main

import (
	"SGin/sgin6"
	"log"
	"net/http"
	"time"
)

func onlyForV2() sgin6.HandlerFunc {
	return func(c *sgin6.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		//c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StateCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := sgin6.New()
	r.Use(sgin6.Logger()) // global midlleware
	r.GET("/", func(c *sgin6.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *sgin6.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
