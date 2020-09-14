package middleware

import (
	"reflect"
	"testing"
)

func TestChain(t *testing.T) {
	mm := Chain()

	type args struct {
		mw Middleware
	}
	tests := []struct {
		name string
		args args
		want Middleware
	}{
		{
			name: "Chain Middleware",
			args: args{mw: mm},
			want: mm,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chain(tt.args.mw); reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chain() = %v, want %v", got, tt.want)
			}
		})
	}
}
