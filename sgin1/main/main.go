package main

import (
	"SGin/sgin1"
	"net/http"
)

func main() {
	c := new(sgin1.Engine)
	http.ListenAndServe(":6600", c)
}
