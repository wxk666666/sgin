package main

import (
	"SGin/sgin3"
	"net/http"
)

func main() {
	e := sgin3.New()
	e.Get("/", func(c *sgin3.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	e.Get("/hello", func(c *sgin3.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	e.Post("/login", func(c *sgin3.Context) {
		c.JSON(http.StatusOK, sgin3.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	e.Run(":6600")
}
