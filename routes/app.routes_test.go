package routes

import (
	"reflect"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestGetRouter(t *testing.T) {
	tests := []struct {
		name string
		want *httprouter.Router
	}{
		{
			name: "Get Router",
			want: httprouter.New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := GetRouter(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
