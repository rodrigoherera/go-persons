package middleware

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

//CORS general middleware
func CORS(next httprouter.Handle) httprouter.Handle {
	return (func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		AllowOrigin := os.Getenv("AllowOrigin")
		AllowMethods := os.Getenv("AllowMethods")
		AllowHeaders := os.Getenv("AllowOrigin")

		w.Header().Set("Access-Control-Allow-Origin", AllowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", AllowMethods)
		w.Header().Set("Access-Control-Allow-Headers", AllowHeaders)
		if r.Method == "OPTIONS" {
			return
		}
		next(w, r, ps)
		return
	})
}
