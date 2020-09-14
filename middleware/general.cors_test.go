package middleware

import (
	"reflect"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestCORS(t *testing.T) {
	var ss httprouter.Handle

	type args struct {
		next httprouter.Handle
	}
	tests := []struct {
		name string
		args args
		want httprouter.Handle
	}{
		{
			name: "Add CORS",
			args: args{next: ss},
			want: ss,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CORS(tt.args.next); reflect.DeepEqual(got, tt.want) {
				t.Errorf("CORS() = %v, want %v", got, tt.want)
			}
		})
	}
}
