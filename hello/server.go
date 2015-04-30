package main

import (
    "github.com/go-martini/martini"
    "github.com/llun/martini-amber"
)

func main() {
    server := martini.Classic()
    server.Use(martini_amber.Renderer(map[string]string{}))
    server.Get("/", func() string {
        return "<h1>Hello</h1>"
    })

    server.Get("/hello", func(r martini_amber.Render) {
        r.Amber(200, "hello", map[string]interface{}{
            "title": "Hello World",
        })
    })

    server.Get("/:who", func(args martini.Params) string {
        return "<h1>Hello " + args["who"] + "</h1>"
    })

    server.Run()
}

