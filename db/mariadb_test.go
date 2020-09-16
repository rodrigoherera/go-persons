package db

import (
	"go-persons/models"
	"net/http"
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
			want:    http.StatusCreated,
			wantErr: false,
		},
		{
			name: "Add Person - BEGIN ERROR",
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
			query:   "",
			want:    http.StatusInternalServerError,
			wantErr: true,
		},
		{
			name: "Add Person - CREATE ERROR",
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
			query:   "",
			want:    http.StatusInternalServerError,
			wantErr: true,
		},
		{
			name: "Add Person - COMMIT ERROR",
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
			query:   "",
			want:    http.StatusInternalServerError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock = mockDb()

			if !tt.wantErr {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(tt.query)).
					WithArgs(tt.args.person.Name, tt.args.person.LastName, tt.args.person.Age, tt.args.person.Dni, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			} else {
				if tt.name == "Add Person - BEGIN ERROR" {
					mock.ExpectExec(regexp.QuoteMeta(tt.query)).
						WithArgs(tt.args.person.Name, tt.args.person.LastName, tt.args.person.Age, tt.args.person.Dni, sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnResult(sqlmock.NewResult(1, 1))
					mock.ExpectRollback()
				}
				if tt.name == "Add Person - CREATE ERROR" {
					mock.ExpectBegin()
					mock.ExpectExec(regexp.QuoteMeta(tt.query)).
						WithArgs(tt.args.person.LastName, tt.args.person.Age, tt.args.person.Dni, sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnResult(sqlmock.NewResult(1, 1))
					mock.ExpectRollback()
				}
				if tt.name == "Add Person - COMMIT ERROR" {
					mock.ExpectBegin()
					mock.ExpectExec(regexp.QuoteMeta(tt.query)).
						WithArgs(tt.args.person.Name, tt.args.person.LastName, tt.args.person.Age, tt.args.person.Dni, sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnResult(sqlmock.NewResult(1, 1))
					mock.ExpectRollback()
				}
			}

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
			want:    http.StatusCreated,
			wantErr: false,
		}, {
			name: "Add User - BEGIN ERROR",
			args: args{
				&models.User{
					Email:    "test@test.com",
					Password: "1234567",
				},
			},
			query:   "",
			want:    http.StatusInternalServerError,
			wantErr: true,
		},
		{
			name: "Add User - CREATE ERROR",
			args: args{
				&models.User{
					Email:    "test@test.com",
					Password: "1234567",
				},
			},
			query:   "",
			want:    http.StatusInternalServerError,
			wantErr: true,
		},
		{
			name: "Add User - COMMIT ERROR",
			args: args{
				&models.User{
					Email:    "test@test.com",
					Password: "1234567",
				},
			},
			query:   "",
			want:    http.StatusInternalServerError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock = mockDb()

			if !tt.wantErr {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(tt.query)).
					WithArgs(tt.args.user.Email, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			} else {
				if tt.name == "Add User - BEGIN ERROR" {
					mock.ExpectExec(regexp.QuoteMeta(tt.query)).
						WithArgs(tt.args.user.Email, sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnResult(sqlmock.NewResult(1, 1))
					mock.ExpectRollback()
				}
				if tt.name == "Add User - CREATE ERROR" {
					mock.ExpectBegin()
					mock.ExpectExec(regexp.QuoteMeta(tt.query)).
						WithArgs(tt.args.user.Email, sqlmock.AnyArg()).
						WillReturnResult(sqlmock.NewResult(1, 1))
					mock.ExpectRollback()
				}
				if tt.name == "Add User - COMMIT ERROR" {
					mock.ExpectBegin()
					mock.ExpectExec(regexp.QuoteMeta(tt.query)).
						WithArgs(tt.args.user.Email, sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnResult(sqlmock.NewResult(1, 1))
					mock.ExpectRollback()
				}
			}

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
		{
			name:    "Get Person - PARAM ERROR",
			query:   "SELECT id, name, lastname, age, dni FROM `people` WHERE (id = ?)",
			wantErr: true,
		},
		{
			name:    "Get Person - SELECT ERROR",
			args:    args{id: "1"},
			query:   "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock = mockDb()

			rows := sqlmock.
				NewRows([]string{"id", "name", "lastname", "age", "dni"}).
				AddRow(p.ID, p.Name, p.LastName, p.Age, p.Dni)

			mock.ExpectQuery(regexp.QuoteMeta(tt.query)).
				WithArgs(p.ID).WillReturnRows(rows)

			if tt.name == "Get Person - SELECT ERROR" {
				Client.Close()
			}
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
			wantErr: false,
		},
		{
			name:    "Get All Person - SELECT ERROR",
			query:   "",
			want:    []models.Person{},
			wantErr: true,
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

			if tt.wantErr {
				Client.Close()
			}

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

func TestDeletePerson(t *testing.T) {
	type args struct {
		person models.Person
	}
	tests := []struct {
		name    string
		args    args
		query   string
		want    int
		wantErr bool
	}{
		{
			name: "Delete Person",
			args: args{
				models.Person{
					Name:        "Test",
					LastName:    "Test",
					Age:         20,
					Dni:         1234567,
					CreatedAt:   time.Now(),
					ProcessedAt: time.Now(),
				},
			},
			query:   "DELETE FROM `people`",
			want:    http.StatusOK,
			wantErr: false,
		},
		{
			name: "Delete Person - BEGIN ERROR",
			args: args{
				models.Person{
					Name:        "Test",
					LastName:    "Test",
					Age:         20,
					Dni:         1234567,
					CreatedAt:   time.Now(),
					ProcessedAt: time.Now(),
				},
			},
			query:   "DELETE FROM `people`",
			want:    http.StatusInternalServerError,
			wantErr: true,
		},
		{
			name: "Delete Person - CREATE ERROR",
			args: args{
				models.Person{
					Name:        "Test",
					LastName:    "Test",
					Age:         20,
					Dni:         1234567,
					CreatedAt:   time.Now(),
					ProcessedAt: time.Now(),
				},
			},
			query:   "DELETE FROM `people`",
			want:    http.StatusInternalServerError,
			wantErr: true,
		},
		{
			name: "Delete Person - COMMIT ERROR",
			args: args{
				models.Person{
					Name:        "Test",
					LastName:    "Test",
					Age:         20,
					Dni:         1234567,
					CreatedAt:   time.Now(),
					ProcessedAt: time.Now(),
				},
			},
			query:   "DELETE FROM `people`",
			want:    http.StatusInternalServerError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := mockDb()

			if !tt.wantErr {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(tt.query)).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			} else {
				if tt.name == "Delete Person - BEGIN ERROR" {
					mock.ExpectExec(regexp.QuoteMeta(tt.query)).
						WillReturnResult(sqlmock.NewResult(1, 1))
					mock.ExpectRollback()
				}
				if tt.name == "Delete Person - CREATE ERROR" {
					mock.ExpectBegin()
					mock.ExpectExec(regexp.QuoteMeta(tt.query)).
						WithArgs(sqlmock.AnyArg()).
						WillReturnResult(sqlmock.NewResult(1, 1))
					mock.ExpectRollback()
				}
				if tt.name == "Delete Person - COMMIT ERROR" {
					mock.ExpectBegin()
					mock.ExpectExec(regexp.QuoteMeta(tt.query)).
						WillReturnResult(sqlmock.NewResult(1, 1))
					mock.ExpectRollback()
				}
			}

			got, err := DeletePerson(tt.args.person)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeletePerson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeletePerson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckUserExistence(t *testing.T) {
	type args struct {
		u *models.User
	}
	tests := []struct {
		name    string
		query   string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "Check User Existence",
			query:   "SELECT id, password FROM `users` WHERE `users`.`id` = ? AND ((email = ?))",
			args:    args{u: &models.User{ID: 1, Email: "test@test.com", Password: "12345678"}},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Check User Existence - SELECT ERROR",
			query:   "SELECT id, password FROM `users` WHERE `users`.`id` = ? AND ((email = ?))",
			args:    args{u: &models.User{ID: 1, Email: "test@test.com", Password: "12345678"}},
			want:    false,
			wantErr: true,
		},
		{
			name:    "Check User Existence - ID ERROR",
			query:   "SELECT id, password FROM `users` WHERE (email = ?)",
			args:    args{u: &models.User{ID: 0, Email: "test@test.com", Password: "12345678"}},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := mockDb()

			rows := sqlmock.
				NewRows([]string{"id", "password"}).
				AddRow(tt.args.u.ID, tt.args.u.Password)

			if tt.name == "Check User Existence - ID ERROR" {
				mock.ExpectQuery(regexp.QuoteMeta(tt.query)).
					WithArgs(tt.args.u.Email).WillReturnRows(rows)
			} else {
				mock.ExpectQuery(regexp.QuoteMeta(tt.query)).
					WithArgs(tt.args.u.ID, tt.args.u.Email).WillReturnRows(rows)
			}

			if tt.name == "Check User Existence - SELECT ERROR" {
				Client.Close()
			}

			got, err := CheckUserExistence(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckUserExistence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckUserExistence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdatePerson(t *testing.T) {
	type args struct {
		person    models.Person
		newPerson models.Person
	}
	tests := []struct {
		name        string
		querySelect string
		queryUpdate string
		args        args
		want        int
		wantErr     bool
	}{
		{
			name:        "Update Person",
			querySelect: "SELECT * FROM `people` ORDER BY `people`.`id` ASC LIMIT 1",
			queryUpdate: "INSERT INTO `people` (`name`,`lastname`,`age`,`dni`,`created_at`,`processed_at`) VALUES (?,?,?,?,?,?)",
			args: args{
				person: models.Person{
					Name:        "Test",
					LastName:    "Test",
					Age:         20,
					Dni:         1234567,
					CreatedAt:   time.Now(),
					ProcessedAt: time.Now(),
				},
				newPerson: models.Person{
					Name:        "Test",
					LastName:    "Test",
					Age:         20,
					Dni:         1234567,
					CreatedAt:   time.Now(),
					ProcessedAt: time.Now(),
				},
			},
			want:    200,
			wantErr: false,
		},
		{
			name:        "Update Person - BEGIN ERROR",
			querySelect: "SELECT * FROM `people` ORDER BY `people`.`id` ASC LIMIT 1",
			queryUpdate: "INSERT INTO `people` (`name`,`lastname`,`age`,`dni`,`created_at`,`processed_at`) VALUES (?,?,?,?,?,?)",
			args: args{
				person: models.Person{
					Name:        "Test",
					LastName:    "Test",
					Age:         20,
					Dni:         1234567,
					CreatedAt:   time.Now(),
					ProcessedAt: time.Now(),
				},
				newPerson: models.Person{
					Name:        "Test",
					LastName:    "Test",
					Age:         20,
					Dni:         1234567,
					CreatedAt:   time.Now(),
					ProcessedAt: time.Now(),
				},
			},
			want:    500,
			wantErr: true,
		},
		{
			name:        "Update Person - SAVE ERROR",
			querySelect: "SELECT * FROM `people` ORDER BY `people`.`id` ASC LIMIT 1",
			queryUpdate: "INSERT INTO `people` (`name`,`lastname`,`age`,`dni`,`created_at`,`processed_at`) VALUES (?,?,?,?,?,?)",
			args: args{
				person: models.Person{
					Name:        "Test",
					LastName:    "Test",
					Age:         20,
					Dni:         1234567,
					CreatedAt:   time.Now(),
					ProcessedAt: time.Now(),
				},
				newPerson: models.Person{
					Name:        "Test",
					LastName:    "Test",
					Age:         20,
					Dni:         1234567,
					CreatedAt:   time.Now(),
					ProcessedAt: time.Now(),
				},
			},
			want:    500,
			wantErr: true,
		},
		{
			name:        "Update Person - COMMIT ERROR",
			querySelect: "SELECT * FROM `people` ORDER BY `people`.`id` ASC LIMIT 1",
			queryUpdate: "INSERT INTO `people` (`name`,`lastname`,`age`,`dni`,`created_at`,`processed_at`) VALUES (?,?,?,?,?,?)",
			args: args{
				person: models.Person{
					Name:        "Test",
					LastName:    "Test",
					Age:         20,
					Dni:         1234567,
					CreatedAt:   time.Now(),
					ProcessedAt: time.Now(),
				},
				newPerson: models.Person{
					Name:        "Test",
					LastName:    "Test",
					Age:         20,
					Dni:         1234567,
					CreatedAt:   time.Now(),
					ProcessedAt: time.Now(),
				},
			},
			want:    500,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := mockDb()
			rows := sqlmock.
				NewRows([]string{"id"}).
				AddRow(tt.args.newPerson.ID)

			if tt.name != "Update Person - BEGIN ERROR" {
				mock.ExpectBegin()
			}
			mock.ExpectQuery(regexp.QuoteMeta(tt.querySelect)).
				WillReturnRows(rows)

			if tt.name == "Update Person - SAVE ERROR" {
				mock.ExpectExec(regexp.QuoteMeta(tt.queryUpdate)).
					WithArgs(sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				mock.ExpectExec(regexp.QuoteMeta(tt.queryUpdate)).
					WillReturnResult(sqlmock.NewResult(1, 1))
			}

			if tt.name != "Update Person - COMMIT ERROR" {
				mock.ExpectCommit()
			} else {
				mock.ExpectRollback()
			}

			got, err := UpdatePerson(tt.args.person, tt.args.newPerson)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdatePerson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdatePerson() = %v, want %v", got, tt.want)
			}
		})
	}
}
