package sgin3

import (
	"fmt"
	"strings"
)

type HandlerFunc func(c *Context)
type router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}
func (r *router) addRoute(m, p string, hf HandlerFunc) {
	var builder strings.Builder

	builder.WriteString(m)
	builder.WriteString("-")
	builder.WriteString(p)

	r.handlers[builder.String()] = hf

}
func (r *router) handle(c *Context) {
	var builder strings.Builder
	builder.WriteString(c.Req.Method)
	builder.WriteString("-")
	builder.WriteString(c.Req.URL.Path)

	if hf, ok := r.handlers[builder.String()]; !ok {
		fmt.Fprintf(c.Writer, "404 NOT FOUND:%s %s\n", c.Req.Method, c.Req.URL)
	} else {
		hf(c)
	}

}
