package routes

import (
	ctrl "go-persons/controller"
	"go-persons/middleware"

	"github.com/julienschmidt/httprouter"
)

// GetRouter returns a composed router
func GetRouter() *httprouter.Router {

	MDLW := middleware.Chain(middleware.JwtAuthentication, middleware.CORS, middleware.JSONHeader)

	router := httprouter.New()

	//TODO - add exception to the middleware
	router.GET("/", ctrl.IndexControlller)
	router.GET("/v1/login/:id", ctrl.Login)
	router.POST("/v1/person", ctrl.AddPerson)

	router.GET("/v1/person", MDLW(ctrl.GetAllPerson))
	router.GET("/v1/person/:id", MDLW(ctrl.GetPerson))
	router.DELETE("/v1/person/:id", MDLW(ctrl.DeletePerson))
	router.PUT("/v1/person/:id", MDLW(ctrl.UpdatePerson))
	return router
}
