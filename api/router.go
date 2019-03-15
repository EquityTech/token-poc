package api

import (
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/gorilla/mux"
	"github.com/ssb4/token-poc/service"
)

// Controller
var controller = &Controller{
	TokenService: service.TokenService{},
}

// Route defines a standard REST route and corresponding handler
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

// Routes represents a slice of routes
type Routes []Route

var routes = Routes{
	Route{
		"CreateToken",
		"POST",
		"/tokens",
		controller.CreateToken,
	},
}

// NewRouter creates a new mux router and adds routes
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// add better formatting for printed routes
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	for _, route := range routes {
		var handler http.Handler
		fmt.Fprintf(w, "%s\t\t%s %s\n", route.Name, route.Method, route.Path)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(handler)
	}
	fmt.Fprintln(w)
	w.Flush()

	return router
}
