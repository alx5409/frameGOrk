package router

import (
	"log"
	"net/http"
	"strings"

	"github.com/alx5409/frameGOrk/middleware"
	"github.com/alx5409/frameGOrk/route"
	"github.com/alx5409/frameGOrk/utils"
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

func allowedMethodsForPath(path route.Path, routes []*route.Route) []string {
	seen := map[string]struct{}{}
	methods := []string{}
	for _, rt := range routes {
		if rt.Path == path {
			m := string(rt.Method)
			if _, ok := seen[m]; !ok {
				seen[m] = struct{}{}
				methods = append(methods, m)
			}
		}
	}
	return methods
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := route.Method(utils.NormalizeMethod(req.Method))
	path := route.Path(req.URL.Path)

	// Checks if the method is correct
	handler, found := HTTPDispatch(method, path, r.routes)
	if found {
		if err := handler(w, req); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	allowed := allowedMethodsForPath(path, r.routes)
	if len(allowed) > 0 {
		w.Header().Set("Allow", strings.Join(allowed, ", "))
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	log.Println("Error:", method, path, "was not found")
	http.NotFound(w, req)
}
