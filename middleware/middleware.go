package middleware

import (
	"github.com/alx5409/frameGOrk/route"
)

// Defines the function signature of a middleware function used by this framework
type Middleware func(route.HandlerFunc) route.HandlerFunc

// Applies middlewares to a final handler in reverse order so that exection happens in the same order middlewares are
// passed in.
func Chain(final route.HandlerFunc, middlewares ...Middleware) route.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		final = middlewares[i](final)
	}
	return final
}
