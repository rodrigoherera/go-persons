package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Middleware provides a convenient mechanism for filtering HTTP requests
// entering the application. It returns a new handler which may perform various
// operations and should finish by calling the next HTTP handler.
type Middleware func(next httprouter.Handle) httprouter.Handle

// Chain provides syntactic sugar to create a new middleware
// which will be the result of chaining the ones received as parameters.
func Chain(mw ...Middleware) Middleware {
	return func(final httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			last := final
			for i := len(mw) - 1; i >= 0; i-- {
				last = mw[i](last)
			}
			last(w, r, ps)
			return
		}
	}
}
