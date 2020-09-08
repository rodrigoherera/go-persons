package controller

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	//AppName project name
	AppName = "Person API"
	//AppVersion Project version
	AppVersion = "1.1"
)

//IndexControlller index controller
func IndexControlller(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	jsonOutput, _ := json.Marshal(struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}{AppName, AppVersion})
	w.Write(jsonOutput)
}
