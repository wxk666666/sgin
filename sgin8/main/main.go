package main

import (
	"SGin/sgin8"
	"SGin/sgin8/middleware"
	"net/http"
)

func main() {
	//eg1:
	r := sgin8.Default()
	config1 := middleware.DefaultCorsConfig()
	cors1 := middleware.Cors{}
	cors1.SetCorsConfig(config1)
	r.Use(cors1.Apply())
	//eg2
	r.Use(middleware.DefaultCorsConfig().Build())
	//eg3
	config3 := middleware.CorsConfig{}
	config3.SetAccessControlMaxAge("200000").SetAccessControlAllowCredentials(true).AddMethods("GET", "POST")
	r.Use(config3.Build())
	//eg4
	config4 := &middleware.CorsConfig{}
	config4.SetAccessControlMaxAge("200000").SetAccessControlAllowCredentials(true).AddMethods("GET", "POST")
	cors4 := middleware.Cors{}
	cors4.SetCorsConfig(config4)
	cors4.Apply()

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
