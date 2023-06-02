package sgin2

import (
	"fmt"
	"net/http"
	"strings"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)
type Engine struct {
	Router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{make(map[string]HandlerFunc)}
}
func (e *Engine) addRoute(m, p string, hf HandlerFunc) {
	var builder strings.Builder

	builder.WriteString(m)
	builder.WriteString("-")
	builder.WriteString(p)

	e.Router[builder.String()] = hf
}
func (e *Engine) Get(p string, hf HandlerFunc) {
	e.addRoute("GET", p, hf)
}
func (e *Engine) Post(p string, hf HandlerFunc) {
	e.addRoute("POST", p, hf)
}
func (e *Engine) Delete(p string, hf HandlerFunc) {
	e.addRoute("DELETE", p, hf)
}
func (e *Engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var builder strings.Builder
	builder.WriteString(r.Method)
	builder.WriteString("-")
	builder.WriteString(r.URL.Path)

	if hf, ok := e.Router[builder.String()]; !ok {
		fmt.Fprintf(w, "404 NOT FOUND:%s %s\n", r.Method, r.URL)
	} else {
		hf(w, r)
	}

}
