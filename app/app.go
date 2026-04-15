package app

import (
	"fmt"
	"net/http"
)

var defaultPort = "8000"

// Contains the metadata and runtime configuration for the framework app.
type App struct {
	Name    string
	Version string
	Addr    string
}

// Starts boots an HTTP server using the provided handler.
// It returns an error if the server fails to start or stops unexpectedly.
func (a *App) Start(handler http.Handler) error {
	if handler == nil {
		return fmt.Errorf("handler cannot be nil")
	}

	if a.Addr == "" {
		a.Addr = ":" + defaultPort
	}
	fmt.Printf("Starting %s version %s on %s\n", a.Name, a.Version, a.Addr)
	return http.ListenAndServe(a.Addr, handler)
}

// Creates a new App instance with name, version and listen address.
func NewApp(name, version, addr string) *App {
	return &App{
		Name:    name,
		Version: version,
		Addr:    addr,
	}
}
