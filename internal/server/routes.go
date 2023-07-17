package server

import (
	"log"

	"github.com/julienschmidt/httprouter"
)

var routes map[string]Route

func init() {
	routes = map[string]Route{
		// Add Routes Here
		"/": {
			method:  "GET",
			handler: homePage,
		},
		"/posts": {
			method:  "GET",
			handler: getPosts,
		},
		"/createpost": {
			method:  "POST",
			handler: createPost,
		},
		"/createuser": {
			method:  "GET",
			handler: createUser,
		},
	}
}
func createRoutes(router *httprouter.Router) {
	for path, route := range routes {

		if route.method == "GET" {
			log.Printf("Registered GET route: %s", path)
			router.GET(path, route.handler)
		}
		if route.method == "POST" {
			log.Printf("Registered POST route: %s", path)
			router.POST(path, route.handler)
		}
	}
}

type Route struct {
	method  string
	handler httprouter.Handle
}
