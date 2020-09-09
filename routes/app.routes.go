package routes

import (
	ctrl "go-persons/controller"
	"go-persons/middleware"
	"os"

	"github.com/newrelic/go-agent/v3/integrations/nrhttprouter"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

// GetRouter returns a composed router
func GetRouter() *nrhttprouter.Router {

	MDLW := middleware.Chain(middleware.JwtAuthentication, middleware.CORS, middleware.JSONHeader)

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("httprouter App"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigDebugLogger(os.Stdout),
	)

	if err != nil {
		panic(err)
	}

	router := nrhttprouter.New(app)

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
