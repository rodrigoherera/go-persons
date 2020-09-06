package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

//Data message data
type Data struct {
	W       http.ResponseWriter
	Message string
	Status  int
}

//Set Entrypoint to ERR
func Set(w http.ResponseWriter, message string, status int) *Data {
	return &Data{
		W:       w,
		Message: message,
		Status:  status,
	}
}

//ReturnJSON Return a JSON message
func (d *Data) ReturnJSON() error {

	jsonOutput, _ := json.Marshal(struct {
		Message string `json:"message"`
	}{d.Message})
	if len(http.StatusText(d.Status)) < 1 {
		return errors.New("Unknow Status")
	}
	d.W.WriteHeader(d.Status)
	d.W.Write(jsonOutput)
	return nil
}

//Return a JSON message
func (d *Data) Return() error {
	d.W.WriteHeader(d.Status)
	d.W.Write([]byte(d.Message))
	return nil
}
