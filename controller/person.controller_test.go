package controller

import (
	"bytes"
	"encoding/json"
	"go-persons/db"
	"go-persons/models"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/julienschmidt/httprouter"
)

var (
	ADDPERSONROUTE    = "/v1/person"
	GETALLPERSONROUTE = "/v1/person"
	GETPERSONROUTE    = "/v1/person/:id"
	UPDATEPERSONROUTE = "/v1/person/:id"
	DELETEPERSONROUTE = "/v1/person/:id"
)

func TestGetPerson(t *testing.T) {
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
			name:     "GET PERSON - Record not found",
			route:    GETPERSONROUTE,
			method:   "GET",
			query:    "SELECT id, name, lastname, age, dni FROM `people` WHERE (id = ?)",
			expected: http.StatusNotFound,
			args:     false,
		},
		{
			name:     "GET PERSON",
			route:    GETPERSONROUTE,
			method:   "GET",
			query:    "SELECT id, name, lastname, age, dni FROM `people` WHERE (id = ?)",
			expected: http.StatusOK,
			args:     true,
		},
	}
	mock = mockDb()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req, err := http.NewRequest(tt.method, "/v1/person/1", nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			router := httprouter.New()

			if tt.args {
				p := models.Person{ID: 1, Name: "test", LastName: "test", Age: 22, Dni: 1234}
				rows := sqlmock.
					NewRows([]string{"id", "name", "lastname", "age", "dni"}).
					AddRow(p.ID, p.Name, p.LastName, p.Age, p.Dni)

				mock.ExpectQuery(regexp.QuoteMeta(tt.query)).
					WithArgs(p.ID).WillReturnRows(rows)
			} else {
				mock.ExpectQuery(regexp.QuoteMeta(tt.query)).
					WithArgs(1).WillReturnRows(sqlmock.NewRows(nil))
			}

			router.GET(tt.route, GetPerson)
			router.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.expected {
				t.Errorf("Wrong status %v", status)
			}
		})
	}
}

func TestAddPerson(t *testing.T) {
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
			name:     "ADD PERSON",
			route:    ADDPERSONROUTE,
			method:   "POST",
			query:    "INSERT INTO `people`",
			expected: http.StatusCreated,
		},
	}
	mock = mockDb()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			person := models.Person{
				Name:        "Test",
				LastName:    "Test",
				Age:         20,
				Dni:         1234567,
				CreatedAt:   time.Now(),
				ProcessedAt: time.Now(),
			}
			requestBody, err := json.Marshal(person)
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
				WithArgs(person.Name, person.LastName, person.Age, person.Dni, sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			router.POST(tt.route, AddPerson)
			router.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.expected {
				t.Errorf("Wrong status %v", status)
			}
		})
	}
}

func TestGetAllPerson(t *testing.T) {
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
			name:     "GET ALL PERSON",
			route:    GETALLPERSONROUTE,
			method:   "GET",
			query:    "SELECT id, name, lastname, age, dni FROM `people`",
			args:     true,
			expected: http.StatusOK,
		},
		{
			name:     "GET ALL PERSON - Not Found",
			route:    GETALLPERSONROUTE,
			method:   "GET",
			query:    "SELECT id, name, lastname, age, dni FROM `people`",
			args:     false,
			expected: http.StatusNotFound,
		},
	}
	mock = mockDb()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.route, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			router := httprouter.New()
			if tt.args {
				p := models.Person{ID: 1, Name: "test", LastName: "test", Age: 22, Dni: 1234}
				rows := sqlmock.
					NewRows([]string{"id", "name", "lastname", "age", "dni"}).
					AddRow(p.ID, p.Name, p.LastName, p.Age, p.Dni)

				mock.ExpectQuery(regexp.QuoteMeta(tt.query)).
					WillReturnRows(rows)
			} else {
				mock.ExpectQuery(regexp.QuoteMeta(tt.query)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows(nil))
			}
			router.GET(tt.route, GetAllPerson)
			router.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.expected {
				t.Errorf("Wrong status %v", status)
			}
		})
	}
}

/* func TestUpdatePerson(t *testing.T) {
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
			name:     "UPDATE PERSON",
			route:    UPDATEPERSONROUTE,
			method:   "PUT",
			query:    "UPDATE `people` SET `name` = ?, `lastname` = ?, `age` = ?, `dni` = ?, `processed_at` = ? WHERE `people`.`id` = ?'",
			args:     true,
			expected: http.StatusOK,
		},
	}
	mock = mockDb()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			person := models.Person{
				Name:        "Test",
				LastName:    "Test",
				Age:         20,
				Dni:         1234567,
				CreatedAt:   time.Now(),
				ProcessedAt: time.Now(),
			}
			requestBody, err := json.Marshal(person)
			if err != nil {
				panic(err)
			}
			req, err := http.NewRequest(tt.method, "/v1/person/1", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			router := httprouter.New()
			if tt.args {
				p := models.Person{ID: 1, Name: "test", LastName: "test", Age: 22, Dni: 1234}
				rows := sqlmock.
					NewRows([]string{"id", "name", "lastname", "age", "dni"}).
					AddRow(p.ID, p.Name, p.LastName, p.Age, p.Dni)

				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, lastname, age, dni FROM `people` WHERE (id = ?)")).
					WithArgs(p.ID).WillReturnRows(rows)

				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `people` WHERE `people.id` = ? ORDER BY `people`.`id` ASC LIMIT 1")).
					WithArgs(0).WillReturnRows(sqlmock.NewRows(nil))

				mock.ExpectExec(tt.query).
					WithArgs(p.Name, p.LastName, p.Age, p.Dni, sqlmock.AnyArg(), p.ID).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			} else {
				mock.ExpectQuery(regexp.QuoteMeta(tt.query)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows(nil))
			}
			router.PUT(tt.route, UpdatePerson)
			router.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.expected {
				t.Errorf("Wrong status %v", status)
			}
		})
	}
} */

func mockDb() sqlmock.Sqlmock {
	mockConn, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	conn, err := gorm.Open("mysql", mockConn)
	if err != nil {
		panic(err)
	}
	db.Client = conn
	return mock
}
