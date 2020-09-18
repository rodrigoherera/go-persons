package controller

import (
	"bytes"
	"encoding/json"
	"go-persons/models"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/julienschmidt/httprouter"
)

var (
	ADDUSERROUTE = "/v1/user"
	LOGINROUTE   = "/v1/login"
)

func TestAddUser(t *testing.T) {
	var mock sqlmock.Sqlmock

	tests := []struct {
		name     string
		route    string
		method   string
		query    string
		expected int
		args     bool
	}{
		{
			name:     "ADD USER",
			route:    ADDUSERROUTE,
			method:   "POST",
			query:    "INSERT INTO `users`",
			expected: http.StatusCreated,
		},
	}
	mock = mockDb()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := models.User{
				Email: "test@test.com",
			}
			requestBody, err := json.Marshal(user)
			if err != nil {
				panic(err)
			}
			req, err := http.NewRequest(tt.method, tt.route, bytes.NewBuffer(requestBody))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			router := httprouter.New()

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(tt.query)).
				WithArgs(user.Email, sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			router.POST(tt.route, AddUser)
			router.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.expected {
				t.Errorf("Wrong status %v", status)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	var mock sqlmock.Sqlmock

	tests := []struct {
		name     string
		route    string
		method   string
		query    string
		expected int
		args     bool
	}{
		{
			name:     "LOGIN",
			route:    LOGINROUTE,
			method:   "POST",
			query:    "SELECT id, password FROM `users` WHERE (email = ?)",
			expected: http.StatusCreated,
		},
	}
	mock = mockDb()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := models.User{
				ID:       1,
				Email:    "test@test.com",
				Password: "12345678",
			}
			req, err := http.NewRequest(tt.method, tt.route, nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetBasicAuth(user.Email, user.Password)

			rr := httptest.NewRecorder()
			router := httprouter.New()

			pass, err := models.GenerateHashPassword("12345678")
			if err != nil {
				panic(err)
			}
			rows := sqlmock.
				NewRows([]string{"id", "password"}).
				AddRow(user.ID, pass)

			mock.ExpectQuery(regexp.QuoteMeta(tt.query)).
				WithArgs(user.Email).WillReturnRows(rows)

			router.POST(tt.route, Login)
			router.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.expected {
				t.Errorf("Wrong status %v", status)
			}
		})
	}
}
