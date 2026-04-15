package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8000"

// Contains the metadata and runtime configuration for the framework app.
type App struct {
	Name    string
	Version string
	Addr    string
	Logger  *log.Logger
}

// sets up a default logger if none is provided
func (a *App) initLogger() {
	if a.Logger == nil {
		a.Logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", a.Name), log.LstdFlags)
	}
}

// validates the following app fields: name, version
func (a *App) validateFields() error {
	if a.Name == "" {
		return fmt.Errorf("app name cannot be empty")
	}
	if a.Version == "" {
		return fmt.Errorf("app version cannot be empty")
	}
	return nil
}

// Boots an HTTP server using the provided handler.
// It returns an error if the server fails to start or stops unexpectedly.
func (a *App) Start(handler http.Handler) error {
	if handler == nil {
		return fmt.Errorf("handler cannot be nil")
	}
	err := a.validateFields()
	if err != nil {
		return err
	}

	if a.Addr == "" {
		a.Addr = ":" + defaultPort
	}
	a.initLogger()
	a.Logger.Printf("Starting %s version %s on %s\n", a.Name, a.Version, a.Addr)
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
