package controller

import (
	"encoding/json"
	"fmt"
	"go-persons/db"
	mid "go-persons/middleware"
	"go-persons/models"
	resp "go-persons/response"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
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

	stringID := strconv.Itoa(int(person.ID))
	if result != 201 {
		http.Error(w, err.Error(), 500)
		return
	}
	resp.Set(w, stringID, 201).Return()
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
	result, err := db.UpdatePerson(person, newPerson, db.Client)
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
	result, err := db.DeletePerson(person, db.Client)
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

//Login login a user
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var creds models.Person

	personID := p.ByName("id")
	if personID == "" {
		http.Error(w, "ID es requerido", 400)
		return
	}
	idInt, err := strconv.Atoi(personID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	creds = models.Person{ID: uint(idInt)}

	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &mid.Claims{
		ID: fmt.Sprint(creds.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(mid.JwtKey)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Printf("Internal server error, error: %v", err)
		return
	}

	result := struct {
		Name    string
		Value   string
		Expires string
	}{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime.String(),
	}
	fmt.Printf("Token generated for User %v", claims.ID)
	resStruct, err := json.Marshal(&result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	resp.Set(w, string(resStruct), http.StatusCreated).Return()
}
