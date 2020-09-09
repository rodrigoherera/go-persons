package routes

import (
	"os"
	"reflect"
	"testing"

	"github.com/newrelic/go-agent/v3/integrations/nrhttprouter"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

func TestGetRouter(t *testing.T) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("httprouter App"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigDebugLogger(os.Stdout),
	)
	if err != nil {
		panic(err)
	}
	tests := []struct {
		name string
		want *nrhttprouter.Router
	}{
		{
			name: "Get Router",
			want: nrhttprouter.New(app),
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
