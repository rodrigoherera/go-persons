package models

import (
	"testing"
	"time"
)

func TestUser_tableName(t *testing.T) {
	type fields struct {
		ID        uint
		Email     string
		Password  string
		CreatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Table name",
			fields: fields{ID: 1, Email: "Test@test.com", Password: "12345678", CreatedAt: time.Now()},
			want:   "user",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:        tt.fields.ID,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
			}
			if got := u.tableName(); got != tt.want {
				t.Errorf("User.tableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateHashPassword(t *testing.T) {
	type args struct {
		pass string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Generate Hash Password",
			args:    args{pass: "12345678"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateHashPassword(tt.args.pass)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateHashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == "" {
				t.Errorf("GenerateHashPassword() = %v", got)
			}
		})
	}
}

func TestCompareHashPasswords(t *testing.T) {
	generated, _ := GenerateHashPassword("12345678")
	type args struct {
		queryPass  string
		actualPass string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Compare Hash Passwords",
			args: args{queryPass: "12345678", actualPass: generated},
			want: true,
		},
		{
			name: "Comapre Hash Passwords - Error",
			args: args{queryPass: "12345678", actualPass: ""},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareHashPasswords(tt.args.queryPass, tt.args.actualPass); got != tt.want {
				t.Errorf("CompareHashPasswords() = %v, want %v", got, tt.want)
			}
		})
	}
}
