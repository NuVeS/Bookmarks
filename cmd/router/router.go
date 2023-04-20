package router

import (
	"fmt"
	"net/http"

	"Bookmarks/cmd/logger"

	"github.com/go-chi/chi/v5"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.Method(route.Method, route.Pattern, handler)
	}

	//router.NotFoundHandler = http.HandlerFunc(urls.Route)

	router.Use()
	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{}
