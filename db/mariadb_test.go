package db

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	PERSONTABLE = "person"
)

/* func TestAddPerson(t *testing.T) {
	type args struct {
		person *models.Person
		c      *gorm.DB
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Add new Person",
			args: args{
				person: &models.Person{Name: "Rodrigo", LastName: "Test", Age: 26, Dni: 1234567},
				c:      new(gorm.DB),
			},
			want:    201,
			wantErr: false,
		},
		{
			name: "Error - nil conn",
			args: args{
				person: nil,
				c:      nil,
			},
			want:    500,
			wantErr: true,
		},
		{
			name: "Error - nil person",
			args: args{
				person: nil,
				c:      new(gorm.DB),
			},
			want:    500,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// dbm, mock, err := sqlmock.New()
			// if err != nil {
			// 	t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			// }
			// rows := sqlmock.NewRows([]string{"id", "name", "lastname", "age", "dni", "created_at", "processed_at"}).
			// 	AddRow(1, "TEST", "Test", 30, 1234566, "2017-11-24 16:56:35").
			// 	AddRow(2, "Tot", "Asd", 24, 2134566, "2018-05-15 00:00:00", "2018-05-15 00:00:00")
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

func TestGetPerson(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    models.Person
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
	tests := []struct {
		name    string
		want    []models.Person
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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

func TestUpdatePerson(t *testing.T) {
	type args struct {
		person    models.Person
		newPerson models.Person
		c         *gorm.DB
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdatePerson(tt.args.person, tt.args.newPerson, tt.args.c)
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

func TestDeletePerson(t *testing.T) {
	type args struct {
		person models.Person
		c      *gorm.DB
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeletePerson(tt.args.person, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeletePerson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeletePerson() = %v, want %v", got, tt.want)
			}
		})
	}
} */
