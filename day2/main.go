package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.Get("/html", func(c *gee.Context) {
		c.Html(http.StatusOK, "<h1>Hello World</h1>")
	})

	r.Get("/string", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Req.URL.Path)
	})

	r.Post("/json", func(c *gee.Context) {
		c.Json(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
