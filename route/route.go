package route

import (
	"fmt"
	"net/http"
)

type Method string
type Path string

// Handles the basic CRUD operations for a RESTful API
const (
	POST   Method = "POST"   // Create
	GET    Method = "GET"    // Read
	PATCH  Method = "PATCH"  // Update
	DELETE Method = "DELETE" // Delete
)

// For now uses the standard library
type HandlerFunc func(http.ResponseWriter, *http.Request) error

type Route struct {
	Path    Path
	Method  Method
	Handler HandlerFunc
}

func (m Method) isValid() bool {
	switch m {
	case POST, GET, PATCH, DELETE:
		return true
	default:
		return false
	}
}

func (p Path) isValid() bool {
	return len(p) > 0 && p[0] == '/'
}

func validateHandler(handler HandlerFunc) bool {
	return handler != nil
}

func validateRoute(method Method, path Path, handler HandlerFunc) error {
	if !method.isValid() {
		return fmt.Errorf("invalid HTTP method: %s", method)
	}
	if !path.isValid() {
		return fmt.Errorf("invalid path: %s", path)
	}
	if !validateHandler(handler) {
		return fmt.Errorf("handler function cannot be nil")
	}
	return nil
}

func New(method Method, path Path, handler HandlerFunc) (*Route, error) {
	if err := validateRoute(method, path, handler); err != nil {
		return nil, err
	}

	return &Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}, nil
}
