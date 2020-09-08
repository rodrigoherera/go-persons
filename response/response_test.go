package response

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	rw := new(httptest.ResponseRecorder)
	d := Data{
		W:       rw,
		Message: "test",
		Status:  200,
	}
	type args struct {
		w       http.ResponseWriter
		message string
		status  int
	}
	tests := []struct {
		name string
		args args
		want *Data
	}{
		{
			name: "Set",
			args: args{w: d.W, message: "test", status: 200},
			want: &d,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Set(tt.args.w, tt.args.message, tt.args.status); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_ReturnJSON(t *testing.T) {
	rw := new(httptest.ResponseRecorder)
	d := Data{
		W:       rw,
		Message: "test",
		Status:  200,
	}

	tests := []struct {
		name    string
		d       *Data
		wantErr bool
	}{
		{
			name:    "Return JSON",
			d:       &d,
			wantErr: false,
		},
		{
			name: "Error Return JSON",
			d: &Data{
				W:       rw,
				Message: "error test",
				Status:  1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.ReturnJSON(); (err != nil) != tt.wantErr {
				t.Errorf("Data.ReturnJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestData_Return(t *testing.T) {
	rw := new(httptest.ResponseRecorder)
	d := Data{
		W:       rw,
		Message: "test",
		Status:  200,
	}
	tests := []struct {
		name    string
		d       *Data
		wantErr bool
	}{
		{
			name:    "Return",
			d:       &d,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Return(); (err != nil) != tt.wantErr {
				t.Errorf("Data.Return() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
