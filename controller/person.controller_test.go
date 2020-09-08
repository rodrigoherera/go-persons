package controller

import (
	"go-persons/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	ADDPERSONROUTE = "/v1/person"
	POSTMETHOD     = "POST"
)

func mockDb() {
	mockConn, _, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	conn, err := gorm.Open("mysql", mockConn)
	if err != nil {
		panic(err)
	}
	db.Client = conn
}

/* func TestAddPerson(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
		p httprouter.Params
	}
	tests := []struct {
		name     string
		method   string
		route    string
		expected int
		args     args
	}{
		{
			name:     "Add new person",
			method:   POSTMETHOD,
			route:    ADDPERSONROUTE,
			expected: http.StatusCreated,
			args:     args{},
		},
	}
	mockDb()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			person := models.Person{
				Name:     "Test",
				LastName: "Test",
				Age:      20,
				Dni:      1234567,
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
			router.POST(tt.route, AddPerson)
			router.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.expected {
				t.Errorf("Wrong status")
			}
		})
	}
} */
