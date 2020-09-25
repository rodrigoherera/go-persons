package routes

import (
	ctrl "go-persons/controller"
	"go-persons/middleware"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetRouter returns a composed router
func GetRouter() *httprouter.Router {

	MDLW := middleware.Chain(middleware.JwtAuthentication, middleware.JSONHeader)

	router := httprouter.New()

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {

			header := w.Header()
			header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			header.Set("Access-Control-Allow-Origin", "*")
			header.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	router.GET("/", MDLW(ctrl.IndexControlller))
	router.GET("/v1/person", MDLW(ctrl.GetAllPerson))
	router.GET("/v1/person/:id", MDLW(ctrl.GetPerson))

	router.POST("/v1/login", MDLW(ctrl.Login))
	router.POST("/v1/user", MDLW(ctrl.AddUser))
	router.POST("/v1/person", MDLW(ctrl.AddPerson))

	router.DELETE("/v1/person/:id", MDLW(ctrl.DeletePerson))
	router.PUT("/v1/person/:id", MDLW(ctrl.UpdatePerson))

	return router
}
