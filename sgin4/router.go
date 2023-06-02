package sgin4

import (
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)
type router struct {
	roots    map[string]*TrieNode
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*TrieNode),
		handlers: make(map[string]HandlerFunc)}
}

// Only one * is allowed
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			//遇到‘*’直接截取，使其成为最后一位
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}
func (r *router) addRoute(m, p string, hf HandlerFunc) {
	parts := parsePattern(p)
	var builder strings.Builder

	builder.WriteString(m)
	builder.WriteString("-")
	builder.WriteString(p)
	//每个方法一个根结点
	_, ok := r.roots[m]
	if !ok {
		newNode := new(TrieNode)
		r.roots[m] = newNode
	}
	r.roots[m].insert(p, parts, 0)
	r.handlers[builder.String()] = hf

}

// In the getRoute function, the parameters of the : and * matches are also parsed, and a map is returned.
// For example, /p/go/doc matches to /p/:lang/doc, the parsing result is: {lang: "go"},
// /static/css/1.css matches to /static/*filepath, and the parsing result is {filepath: "css/1.css"}.
func (r *router) getRoute(method string, path string) (*TrieNode, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)
	if n != nil {
		//找到对于前缀树中的路径(可能含':' or '*'的路径)
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index] //:lang/code and c/code -> [lang]=c
			}
			if part[0] == '*' && len(part) > 1 {
				//index之后的就是所替代的
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break // c/*filepath and c/d/1.css -> [filepath]=d/1.css
			}
		}
		return n, params
	}
	return nil, nil
}
func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		f, ok := r.handlers[c.Method+"-"+n.pattern]
		if !ok {
			c.String(http.StatusInternalServerError, "500 Server Error: %s\n", c.Path)
			return
		}
		f(c)

	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}

}
