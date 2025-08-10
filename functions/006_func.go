package main

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

var routeHandlers = map[string]HandlerFunc{}

func registerRoute(path string, handler HandlerFunc) {
	routeHandlers[path] = handler
}

func mainRouter(w http.ResponseWriter, r *http.Request) {
	if handler, ok := routeHandlers[r.URL.Path]; ok {
		handler(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	registerRoute("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from /hello")
	})

	registerRoute("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Service is up and running")
	})

	http.HandleFunc("/", mainRouter)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

/*
Output (server starts and listens)
Server started on :8080
(server runs until terminated)

When visiting URLs:
http://localhost:8080/hello → "Hello from /hello"
http://localhost:8080/status → "Service is up and running"
http://localhost:8080/xyz → 404 Not Found
*/

/*
Code Explanation:
- Purpose: Demonstrate function types, function variables, and HTTP routing with anonymous functions
- HandlerFunc is a type alias for func(http.ResponseWriter, *http.Request)
- routeHandlers map stores path-to-function mappings using the HandlerFunc type
- registerRoute stores anonymous functions in the map for different URL paths
- mainRouter looks up handlers by request path and calls them, or returns 404
- Anonymous functions are defined inline when registering routes
- http.ListenAndServe starts the server and blocks until terminated
*/
