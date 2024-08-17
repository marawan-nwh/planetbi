package handlers

import (
	"errors"

	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type handler struct {
	URI    string
	Method string
	Params string
	Auth   bool
	Role   string
	Do     http.HandlerFunc
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// validate request parameters
	if len(h.Params) > 0 {
		err := validate(r, h.Params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	h.Do(w, r)
}

// Handle registers arrays of handlers
func Handle(router *mux.Router, groups ...[]handler) {
	// To be able to register different handlers for the same uri with the standard http.ServeMux,
	// we don't register the handler directly, we create a top layer handlerFunc
	// that routes the request to the right handler based on the request method.
	var routes = make(map[string]map[string]handler)
	for _, group := range groups {
		for _, h := range group {
			if routes[h.URI] == nil {
				routes[h.URI] = make(map[string]handler)
			}

			method := strings.ToUpper(h.Method)
			if routes[h.URI][method].URI != "" {
				panic("Duplicate Routes: " + method + " " + h.URI)
			}
			routes[h.URI][strings.ToUpper(h.Method)] = h
		}
	}

	for uri, m := range routes {
		router.HandleFunc(uri, routeFunc(uri, m))
	}
}

func routeFunc(uri string, m map[string]handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if m[r.Method].Method != "" {
			m[r.Method].ServeHTTP(w, r)
			return
		}

		for k := range m {
			m[k].ServeHTTP(w, r)
			return
		}

		http.NotFound(w, r)
	}
}

// validate presence of params and checks specified validations
// empty: can be empty
// max=x: can't be more than x
// min=x: can't be less than x
// numeric: should be numerical value
func validate(r *http.Request, prms string) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	params := strings.Split(prms, ",")

	for _, param := range params {
		p := strings.Split(strings.TrimSpace(param), ":")

		if len(p) < 2 { // no rules, just validate presence
			if r.FormValue(p[0]) == "" {
				return errors.New(p[0] + " can't be empty")
			}
			continue
		}

		if r.FormValue(p[0]) == "" && !strings.Contains(p[1], "empty") {
			return errors.New(p[0] + " can't be empty")
		}

		s := strings.Split(p[1], "&")

		for _, c := range s {
			vs := strings.Split(c, "=")
			switch vs[0] {
			case "max":
				val, _ := strconv.Atoi(vs[1])
				prm := strings.TrimSpace(r.FormValue(p[0]))
				if len([]rune(prm)) > val {
					return errors.New(p[0] + " can't be more than " + vs[1] + " characters")
				}
			case "min":
				val, _ := strconv.Atoi(vs[1])
				prm := strings.TrimSpace(r.FormValue(p[0]))
				if len([]rune(prm)) < val {
					return errors.New(p[0] + " can't be less than " + vs[1] + " characters")
				}
			case "numeric":
				_, err := strconv.Atoi(r.FormValue(p[0]))
				if err != nil {
					return errors.New(p[0] + " should be an integer string")
				}

			}
		}

	}
	return nil
}
