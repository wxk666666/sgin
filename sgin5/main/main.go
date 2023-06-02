package main

import (
	"SGin/sgin5"
	"net/http"
)

func main() {
	r := sgin5.New()
	r.GET("/index", func(c *sgin5.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	//This "{ }" is only used to visualize the structure of the route
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *sgin5.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *sgin5.Context) {
			// expect /hello?name=wxk
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{ //This {} only looks at the v2 range
		v2.GET("/hello/:name", func(c *sgin5.Context) {
			// expect /hello/wxk
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *sgin5.Context) {
			c.JSON(http.StatusOK, sgin5.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":6600")
}
