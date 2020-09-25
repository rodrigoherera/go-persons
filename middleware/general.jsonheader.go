package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//JSONHeader general middleware
func JSONHeader(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "aplication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r, ps)
		return
	}
}
