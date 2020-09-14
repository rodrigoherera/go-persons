package middleware

import (
	"fmt"
	"go-persons/models"
	"go-persons/response"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

var JwtKey = []byte("6B833C42246DA36D0DCC912DE9220DE4F6DD146321639B59AC4D1B9BD226A228")

//JwtAuthentication validate the jwt for the incoming request
func JwtAuthentication(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		notAuth := []string{"/v1/login", "/v1/user", "/"}
		requestPath := r.URL.Path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {

			if value == requestPath {
				next(w, r, ps)
				return
			}
		}

		c := r.Header.Get("Authorization")
		if c == "" {
			response.Set(w, "El Token JWT es requerido!", 403).ReturnJSON()
			log.Println("Missing auth token")
			return
		}

		tknStr := strings.TrimPrefix(c, "Bearer ")

		claims := &models.UserClaim{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				// For any other type of error, return a bad request status
				response.Set(w, err.Error(), 400).ReturnJSON()
				log.Printf("JWT - Bad Request, error: %v", err)
				return
			}
			response.Set(w, err.Error(), 400).ReturnJSON()
			log.Println("JWT expired - Has to renew de JWT")
			return
		}
		if !tkn.Valid {
			response.Set(w, "Token no valido", 403).ReturnJSON()
			log.Printf("Unauthorized Request, token no valid")
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		fmt.Printf("User: %v, is login ", claims.Email) //Useful for monitoring
		next(w, r, ps)
	}
}

// //Renew renew the token
// func Renew(w http.ResponseWriter, r *http.Request) {
// 	c := r.Header.Get("Authorization")
// 	if c == "" {
// 		response.Set(w, "El Token JWT es requerido!", 403).ReturnJSON()
// 		log.Println("Missing auth token")
// 		return
// 	}

// 	tknStr := strings.TrimPrefix(c, "Bearer ")

// 	claims := &claims{}

// 	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			// For any other type of error, return a bad request status
// 			response.Set(w, "El Token JWT es requerido!", 403).ReturnJSON()
// 			log.Printf("Unauthorized Request, error: %v", err)
// 			return
// 		}
// 	}
// 	if !tkn.Valid {
// 		// (END) The code uptil this point is the same as the first part of the `Welcome` route

// 		// We ensure that a new token is not issued until enough time has elapsed
// 		// In this case, a new token will only be issued if the old token is within
// 		// 30 seconds of expiry. Otherwise, return a bad request status
// 		if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}

// 		// Now, create a new token for the current use, with a renewed expiration time
// 		expirationTime := time.Now().Add(12 * time.Hour)
// 		claims.ExpiresAt = expirationTime.Unix()
// 		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 		tokenString, err := token.SignedString(jwtKey)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		result := struct {
// 			Name    string
// 			Value   string
// 			Expires string
// 		}{
// 			Name:    "token",
// 			Value:   tokenString,
// 			Expires: expirationTime.String(),
// 		}
// 		fmt.Printf("Token updated for User %v", claims.Email)
// 		res.RespondWithJSON(w, http.StatusCreated, result)
// 	}

// }
