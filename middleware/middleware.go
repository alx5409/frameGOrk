package middleware

import (
	"github.com/alx5409/frameGOrk/route"
)

type Middleware func(route.HandlerFunc) route.HandlerFunc

func Chain(final route.HandlerFunc, middlewares ...Middleware) route.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		final = middlewares[i](final)
	}
	return final
}
