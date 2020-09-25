package controller

import (
	"encoding/json"
	"fmt"
	"go-persons/db"
	mid "go-persons/middleware"
	"go-persons/models"
	resp "go-persons/response"
	"io/ioutil"
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
		resp.Set(w, "Body required", 400).Return()
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.Set(w, err.Error(), 500).Return()
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(data, &user)
	if err != nil {
		resp.Set(w, err.Error(), 500).Return()
		return
	}

	if !strings.Contains(user.Email, "@") {
		resp.Set(w, "Invalid email format", 500).Return()
		return
	}

	if len(user.Password) < 0 {
		resp.Set(w, "Password required", 400).Return()
		return
	}

	hashedPassword, err := models.GenerateHashPassword(user.Password)
	if err != nil {
		resp.Set(w, err.Error(), 500).Return()
		return
	}

	user.Password = hashedPassword

	result, err := db.AddUser(&user)
	if err != nil {
		resp.Set(w, err.Error(), 500).Return()
		return
	}

	stringID := strconv.Itoa(int(user.ID))
	if result != 201 {
		resp.Set(w, err.Error(), 500).Return()
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
		resp.Set(w, err.Error(), 500).Return()
		return
	}
	resp.Set(w, string(userByte), 201).Return()
}

//Login login a user
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	email, pass, ok := r.BasicAuth()
	if !ok {
		resp.Set(w, "Basic Auth requerida", 400).Return()
		return
	}
	if !strings.Contains(email, "@") {
		resp.Set(w, "Invalid email format", 400).Return()
		return
	}

	if len(pass) < 0 {
		resp.Set(w, "Password required", 400).Return()
		return
	}

	u := models.User{Email: email}

	result, err := db.CheckUserExistence(&u)
	if err != nil && err.Error() != "record not found" {
		resp.Set(w, err.Error(), 500).Return()
		return
	}

	if !result {
		resp.Set(w, "User not found, email or password invalid", 404).Return()
		return
	}

	equalPass := models.CompareHashPasswords(pass, u.Password)
	if !equalPass {
		resp.Set(w, "User not found, email or password invalid", 404).Return()
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
		resp.Set(w, err.Error(), 500).Return()
		return
	}

	resultJSON := struct {
		Email   string `json:"email"`
		Name    string `json:"name"`
		Value   string `json:"value"`
		Expires string `json:"expires"`
	}{
		Email:   u.Email,
		Name:    "Bearer token",
		Value:   tokenString,
		Expires: expirationTime.String(),
	}
	fmt.Printf("Token generated for User %v", claims.Email)

	resStruct, err := json.Marshal(&resultJSON)
	if err != nil {
		resp.Set(w, err.Error(), 500).Return()
		return
	}
	resp.Set(w, string(resStruct), http.StatusCreated).Return()
}
