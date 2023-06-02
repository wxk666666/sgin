package sgin3

import (
	"net/http"
)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: NewRouter()}
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
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}
func (e *Engine) addRoute(m, p string, hf HandlerFunc) {
	e.router.addRoute(m, p, hf)
}
