package main

import (
	"fmt"
	"net/http"

	"github.com/alx5409/frameGOrk/route"
	"github.com/alx5409/frameGOrk/router"
)

func main() {
	router := &router.Router{}
	methods := []route.Method{route.GET, route.POST, route.PATCH, route.DELETE}
	hFunc := route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		fmt.Fprintln(w, "Hello, World!")
		return nil
	})

	for _, method := range methods {
		if err := router.Register(method, "/hello", hFunc); err != nil {
			fmt.Printf("Error registering route: %v\n", err)
		}
	}

	routes := router.GetRoutes()
	for _, r := range routes {
		fmt.Printf("Registered route: %s %s\n", r.Method, r.Path)
	}
}
