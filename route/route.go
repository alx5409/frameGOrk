package route

import (
	"fmt"
)

type Route struct {
}

func createRoute() *Route {
	fmt.Println("Creating a new route")
	return &Route{}
}
