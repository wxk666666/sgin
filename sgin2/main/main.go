package main

import (
	"SGin/sgin2"
	"fmt"
	"net/http"
)

func main() {
	engine := sgin2.New()
	engine.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get /")
	})
	engine.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get /Hello")
	})
	engine.Run(":6060")
}
