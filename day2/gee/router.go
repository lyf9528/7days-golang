package gee

import (
	"fmt"
	"log"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	r := &router{
		handlers: make(map[string]HandlerFunc),
	}
	return r
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Req.Method + "-" + c.Req.URL.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		fmt.Fprintf(c.Writer, "404 NOT FOUND: %s\n", key)
	}
}
