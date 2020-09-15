package controller

import (
	"encoding/json"
	"go-persons/db"
	"go-persons/models"
	resp "go-persons/response"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//AddPerson POST controller to add a new person
func AddPerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var person models.Person
	if r.Body == nil {
		http.Error(w, "BODY es requerido", 400)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(data, &person)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	result, err := db.AddPerson(&person)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if result != 201 {
		http.Error(w, err.Error(), 500)
		return
	}

	personByte, err := json.Marshal(person)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	resp.Set(w, string(personByte), 201).Return()
}

//GetPerson get a person by a giving IDE
func GetPerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	personID := p.ByName("id")
	if personID == "" {
		http.Error(w, "ID es requerido", 400)
		return
	}
	person, err := db.GetPerson(personID)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	personJSON, err := json.Marshal(person)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	resp.Set(w, string(personJSON), 200).Return()
}

//GetAllPerson get all persons
func GetAllPerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	person, err := db.GetAllPerson()
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	personJSON, err := json.Marshal(person)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	resp.Set(w, string(personJSON), 200).Return()
}

//UpdatePerson update a person
func UpdatePerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newPerson models.Person

	personID := p.ByName("id")
	if personID == "" {
		http.Error(w, "ID es requerido", 400)
		return
	}

	if r.Body == nil {
		http.Error(w, "BODY es requerido", 400)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(data, &newPerson)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	person, err := db.GetPerson(personID)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	result, err := db.UpdatePerson(person, newPerson)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if result != 200 {
		http.Error(w, err.Error(), 500)
		return
	}
	resp.Set(w, "OK", 200).Return()
}

//DeletePerson delete a person
func DeletePerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	personID := p.ByName("id")
	if personID == "" {
		http.Error(w, "ID es requerido", 400)
		return
	}

	person, err := db.GetPerson(personID)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	result, err := db.DeletePerson(person)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if result != 200 {
		http.Error(w, err.Error(), 500)
		return
	}
	resp.Set(w, "OK", 200).Return()
}
