package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

//User contains user structure
type User struct {
	ID        uint      `gorm:"primary_key"`
	Email     string    `json:"email" gorm:"column:email"`
	Password  string    `json:"password" gorm:"column:password"`
	CreatedAt time.Time `json:"-" gorm:"column:created_at"`
}

//UserClaim used to return JWT AUTH
type UserClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (u *User) tableName() string {
	return "user"
}

//GenerateHashPassword generate a new hash password
func GenerateHashPassword(pass string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

//CompareHashPasswords given 2 passwords return if the hash is equal
func CompareHashPasswords(queryPass, actualPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(actualPass), []byte(queryPass))
	if err != nil { //Password does not match!
		return false
	}
	return true
}
