package router

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

const paramPattern = "([^/]+)"
const uuidPattern = "([a-fA-F0-9-]+)"

func NewRouter() http.Handler {
	return http.HandlerFunc(Serve)
}

func HomeHandler(w http.ResponseWriter, h *http.Request) {
  fmt.Fprint(w, "Home")
}

var routes = []route{
  newRoute("GET", "/", HomeHandler),
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
  return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type route struct {
  method  string
  regex   *regexp.Regexp
  handler http.HandlerFunc
}

func Serve(w http.ResponseWriter, r *http.Request) {
  var allow []string

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
    http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
  }

  http.NotFound(w, r)
}

type ctxKey struct{}

func getFields(r *http.Request, index int) string {
  fields := r.Context().Value(ctxKey{}).([]string)
  return fields[index]
}

