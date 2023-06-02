package main

import (
	"SGin/sgin4"
	"net/http"
)

func main() {
	r := sgin4.New()
	r.Get("/", func(c *sgin4.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.Get("/hello", func(c *sgin4.Context) {
		// expect /hello?name=wxk
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.Get("/hello/:name", func(c *sgin4.Context) {
		// expect /hello/wxk
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.Get("/assets/*filepath", func(c *sgin4.Context) {
		c.JSON(http.StatusOK, sgin4.H{"filepath": c.Param("filepath")})
	})

	r.Run(":6600")
}
