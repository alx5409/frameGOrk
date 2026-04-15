package route

import (
	"fmt"
	"net/http"
)

// HTTP method used for router matching.
type Method string

// URL path pattern for a route.
type Path string

// Handles the basic CRUD operations for a RESTful API.
const (
	POST   Method = "POST"   // Create
	GET    Method = "GET"    // Read
	PATCH  Method = "PATCH"  // Update
	DELETE Method = "DELETE" // Delete
)

// Defines the handler signature of a handler function.
// For now uses the standard library.
type HandlerFunc func(http.ResponseWriter, *http.Request) error

// Defines a single registered endpoint with its method, path and handler function.
type Route struct {
	Path    Path
	Method  Method
	Handler HandlerFunc
}

// checks if the method is a valid method.
func (m Method) isValid() bool {
	switch m {
	case POST, GET, PATCH, DELETE:
		return true
	default:
		return false
	}
}

// checks if the path is a valid path.
func (p Path) isValid() bool {
	return len(p) > 0 && p[0] == '/'
}

// checks if the handler is a valid handler function.
func validateHandler(handler HandlerFunc) bool {
	return handler != nil
}

// validates method, path and handler function before creating a route.
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

// Creates and validates a Route instance.
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
