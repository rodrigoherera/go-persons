package controller

import (
	"encoding/json"
	"fmt"
	"go-persons/db"
	"go-persons/models"
	"go-persons/response"
	resp "go-persons/response"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

type claims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

var jwtKey = []byte("6B833C42246DA36D0DCC912DE9220DE4F6DD146321639B59AC4D1B9BD226A228")

//AddPerson POST controller to add a new person
func AddPerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var person models.Person
	if r.Body == nil {
		//asdsa
		//exploto o retorno error
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//algo
	}
	defer r.Body.Close()

	err = json.Unmarshal(data, &person)
	if err != nil {

	}

	result, err := db.AddPerson(&person, db.Client)
	if err != nil {

	}

	stringID := strconv.Itoa(int(person.ID))
	if result == 201 {
		resp.Set(w, stringID, 200).Return()
		return
	}
	resp.Set(w, "No se pudo insertar usuario", 500).Return()
}

//GetPerson get a person by a giving IDE
func GetPerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	personID := p.ByName("id")
	if personID != "" {
		//algo
	}
	person, err := db.GetPerson(personID)
	if err != nil {

	}
	personJSON, err := json.Marshal(person)
	if err != nil {
		//
	}
	resp.Set(w, string(personJSON), 200).Return()
}

//GetAllPerson get all persons
func GetAllPerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	person, err := db.GetAllPerson()
	if err != nil {

	}
	personJSON, err := json.Marshal(person)
	if err != nil {
		//
	}
	resp.Set(w, string(personJSON), 200).Return()
}

//UpdatePerson update a person
func UpdatePerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newPerson models.Person

	personID := p.ByName("id")
	if personID != "" {
		//algo
	}

	if r.Body == nil {
		//asdsa
		//exploto o retorno error
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//algo
	}
	defer r.Body.Close()

	err = json.Unmarshal(data, &newPerson)
	if err != nil {

	}

	person, err := db.GetPerson(personID)
	if err != nil {

	}
	result, err := db.UpdatePerson(person, newPerson, db.Client)
	if err != nil {

	}
	if result != 200 {
		resp.Set(w, "Error", 500).Return()
		return
	}
	resp.Set(w, "OK", 200).Return()
}

//DeletePerson delete a person
func DeletePerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newPerson models.Person

	personID := p.ByName("id")
	if personID != "" {
		//algo
	}

	if r.Body == nil {
		//asdsa
		//exploto o retorno error
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//algo
	}
	defer r.Body.Close()

	err = json.Unmarshal(data, &newPerson)
	if err != nil {

	}

	person, err := db.GetPerson(personID)
	if err != nil {

	}
	result, err := db.DeletePerson(person, db.Client)
	if err != nil {

	}
	if result != 200 {
		resp.Set(w, "Error", 500).Return()
		return
	}
	resp.Set(w, "OK", 200).Return()
}

//Login asdsa
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var creds models.Person

	personID := p.ByName("id")
	if personID != "" {
		//algo
	}
	idInt, err := strconv.Atoi(personID)
	if err != nil {
		//
	}
	creds = models.Person{ID: uint(idInt)}

	expirationTime := time.Now().Add(12 * time.Hour)
	claims := &claims{
		ID: string(creds.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		response.Set(w, "El Token JWT es requerido!", 403).ReturnJSON()
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
		//
	}
	resp.Set(w, string(resStruct), http.StatusCreated).Return()
}
