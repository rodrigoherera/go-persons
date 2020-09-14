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
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

//AddUser add a new user
func AddUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var user models.User

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

	err = json.Unmarshal(data, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	hashedPassword, err := models.GenerateHashPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	user.Password = hashedPassword

	result, err := db.AddUser(&user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	stringID := strconv.Itoa(int(user.ID))
	if result != 201 {
		http.Error(w, err.Error(), 500)
		return
	}

	userJSON := struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}{
		ID:    stringID,
		Email: user.Email,
	}
	userByte, err := json.Marshal(userJSON)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	resp.Set(w, string(userByte), 201).Return()
}

//Login login a user
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	email, pass, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "Basic Auth requerida", 400)
	}
	if !strings.Contains(email, "@") {
		http.Error(w, "Emal invalido", 400)
		return
	}

	if len(pass) < 0 {
		http.Error(w, "Password requerida", 400)
		return
	}

	u := models.User{Email: email}

	result, err := db.CheckUserExistence(&u)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if !result {
		http.Error(w, err.Error(), 404)
		return
	}

	equalPass := models.CompareHashPasswords(pass, u.Password)
	if !equalPass {
		http.Error(w, "Password incorrecta", 404)
		return
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &models.UserClaim{
		Email: u.Email,
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

	resultJSON := struct {
		Email   string
		Name    string
		Value   string
		Expires string
	}{
		Email:   u.Email,
		Name:    "Bearer token",
		Value:   tokenString,
		Expires: expirationTime.String(),
	}
	fmt.Printf("Token generated for User %v", claims.Email)

	resStruct, err := json.Marshal(&resultJSON)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	resp.Set(w, string(resStruct), http.StatusCreated).Return()
}
