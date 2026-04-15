package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alx5409/frameGOrk/app"
	"github.com/alx5409/frameGOrk/route"
	"github.com/alx5409/frameGOrk/router"
)

func main() {
	appName := "frameGOrk"
	appVersion := "0.0.1"
	appAddress := ":8000"
	mockPath := "/hello"
	handlerFunction := route.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) error {
			fmt.Fprintln(w, "Hello, this is a demo")
			return nil
		})

	// Creates the app
	app := app.NewApp(appName, appVersion, appAddress)

	// Create the router
	r := &router.Router{}

	// Register routes
	methods := []route.Method{route.GET, route.POST, route.PATCH, route.DELETE}

	for _, method := range methods {
		err := r.Register(method, route.Path(mockPath), handlerFunction)
		if err != nil {
			log.Printf("Error registering route: %v\n", err)
		}
	}

	var URL = "localhost" + appAddress + mockPath
	fmt.Println("Server is running on ", URL)

	// Start the server using the router
	err := app.Start(r)
	if err != nil {
		log.Fatalf("Server error: %v\n", err)
	}
}
