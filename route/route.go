package route

import (
	"errors"
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
type HandlerFunc func(http.Request)

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
		errMessage := fmt.Errorf("invalid HTTP method: %s", method)
		return errors.New(errMessage.Error())
	}
	if !path.isValid() {
		errMessage := fmt.Errorf("invalid path: %s", path)
		return errors.New(errMessage.Error())
	}
	if !validateHandler(handler) {
		return errors.New("handler function cannot be nil")
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
