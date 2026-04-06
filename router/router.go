package router

import (
	"github.com/alx5409/frameGOrk/middleware"
	"github.com/alx5409/frameGOrk/route"
)

type Router struct {
	routes []*route.Route
}

func (r *Router) Register(method route.Method, path route.Path, handler route.HandlerFunc, middlewares ...middleware.Middleware) error {
	finalHandler := middleware.Chain(handler, middlewares...)
	newRoute := &route.Route{
		Method:  method,
		Path:    path,
		Handler: finalHandler,
	}
	r.routes = append(r.routes, newRoute)
	return nil
}

func (r *Router) GetRoutes() []*route.Route {
	return r.routes
}

func HTTPDispatch(method route.Method, path route.Path, routes []*route.Route) (route.HandlerFunc, bool) {
	for _, r := range routes {
		if r.Method == method && r.Path == path {
			return r.Handler, true
		}
	}
	return nil, false
}
