package db

import (
	"go-persons/models"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	PERSONTABLE = "person"
)

func TestAddPerson(t *testing.T) {
	var mock sqlmock.Sqlmock
	type args struct {
		person *models.Person
	}
	tests := []struct {
		name    string
		args    args
		query   string
		want    int
		wantErr bool
	}{
		{
			name: "Add Person",
			args: args{
				&models.Person{
					Name:        "Test",
					LastName:    "Test",
					Age:         20,
					Dni:         1234567,
					CreatedAt:   time.Now(),
					ProcessedAt: time.Now(),
				},
			},
			query:   "INSERT INTO `people`",
			want:    201,
			wantErr: false,
		},
	}
	mock = mockDb()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(tt.query)).
				WithArgs(tt.args.person.Name, tt.args.person.LastName, tt.args.person.Age, tt.args.person.Dni, sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			got, err := AddPerson(tt.args.person)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddPerson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddPerson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddUser(t *testing.T) {
	var mock sqlmock.Sqlmock

	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		args    args
		query   string
		want    int
		wantErr bool
	}{
		{
			name: "Add User",
			args: args{
				&models.User{
					Email:    "test@test.com",
					Password: "1234567",
				},
			},
			query:   "INSERT INTO `users`",
			want:    201,
			wantErr: false,
		},
	}
	mock = mockDb()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(tt.query)).
				WithArgs(tt.args.user.Email, sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			got, err := AddUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPerson(t *testing.T) {
	var mock sqlmock.Sqlmock

	p := models.Person{ID: 1, Name: "test", LastName: "test", Age: 22, Dni: 1234}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    models.Person
		query   string
		wantErr bool
	}{
		{
			name:    "Get Person",
			args:    args{id: "1"},
			want:    p,
			query:   "SELECT id, name, lastname, age, dni FROM `people` WHERE (id = ?)",
			wantErr: false,
		},
	}
	mock = mockDb()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows := sqlmock.
				NewRows([]string{"id", "name", "lastname", "age", "dni"}).
				AddRow(p.ID, p.Name, p.LastName, p.Age, p.Dni)

			mock.ExpectQuery(regexp.QuoteMeta(tt.query)).
				WithArgs(p.ID).WillReturnRows(rows)
			got, err := GetPerson(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPerson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPerson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllPerson(t *testing.T) {
	var mock sqlmock.Sqlmock

	tests := []struct {
		name    string
		query   string
		want    []models.Person
		wantErr bool
	}{
		{
			name:  "Get All Person",
			query: "SELECT id, name, lastname, age, dni FROM `people`",
			want: []models.Person{
				models.Person{ID: 1, Name: "test", LastName: "test", Age: 22, Dni: 1234},
				models.Person{ID: 1, Name: "test", LastName: "test", Age: 22, Dni: 1234},
			},
		},
	}

	mock = mockDb()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := models.Person{ID: 1, Name: "test", LastName: "test", Age: 22, Dni: 1234}
			rows := sqlmock.
				NewRows([]string{"id", "name", "lastname", "age", "dni"}).
				AddRow(p.ID, p.Name, p.LastName, p.Age, p.Dni).
				AddRow(p.ID, p.Name, p.LastName, p.Age, p.Dni)

			mock.ExpectQuery(regexp.QuoteMeta(tt.query)).
				WillReturnRows(rows)
			got, err := GetAllPerson()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllPerson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllPerson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func mockDb() sqlmock.Sqlmock {
	mockConn, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	conn, err := gorm.Open("mysql", mockConn)
	if err != nil {
		panic(err)
	}
	Client = conn
	return mock
}
