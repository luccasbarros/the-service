package api

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	data "github.com/luccasbarros/the-service/internal/postgres"
)

// const paramPattern = "([^/]+)"
// const uuidPattern = "([a-fA-F0-9-]+)"

func NewHandler(dal *data.Data) http.Handler {
	appHandler := NewAppHandler(dal)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Serve(w, r, appHandler)
	})
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

func Serve(w http.ResponseWriter, r *http.Request, appHandler *AppHandler) {
	var allow []string

	var routes = []route{
		// users
		newRoute("GET", "/", appHandler.UsersHandler.GetAllUsers),
	}

	// documentation UI
	routes = append(routes, newRoute("GET", "/docs", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../spec/redoc-static.html")
	})))

	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}

			ctx := context.WithValue(r.Context(), ctxKey{}, matches[:1])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}

	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ","))
		RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}

	RespondError(w, http.StatusNotFound, "Not found")
}

type ctxKey struct{}

func GetFields(r *http.Request, index int) string {
	fields := r.Context().Value(ctxKey{}).([]string)
	return fields[index]
}
