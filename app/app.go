package app

import (
	"fmt"
)

type App struct {
	Name    string
	Version string
}

func (a *App) Start() {
	fmt.Printf("Starting %s version %s\n", a.Name, a.Version)
}

func NewApp(name, version string) *App {
	return &App{
		Name:    name,
		Version: version,
	}
}
