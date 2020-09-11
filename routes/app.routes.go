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

	router.GET("/", MDLW(ctrl.IndexControlller))
	router.GET("/v1/login", MDLW(ctrl.Login))
	router.GET("/v1/person", MDLW(ctrl.GetAllPerson))
	router.GET("/v1/person/:id", MDLW(ctrl.GetPerson))

	router.POST("/v1/user", MDLW(ctrl.AddUser))
	router.POST("/v1/person", MDLW(ctrl.AddPerson))

	router.DELETE("/v1/person/:id", MDLW(ctrl.DeletePerson))
	router.PUT("/v1/person/:id", MDLW(ctrl.UpdatePerson))

	return router
}
