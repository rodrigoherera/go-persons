package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func TestJWTValidate(t *testing.T) {
	type args struct {
		jwt string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "JWTValidate",
			args: args{
				jwt: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzZXNzaW9uVG9rZW4iOiIwMGFjNzQ1NmYxM2NiMjhhYmM2YzdjNGVkY2UwOTc4MjRjMmYxNDE2IiwiY29kY2xpIjoxMDU0NiwiY29kX3BhcnRuZXIiOjQ1Niwic29sdWNpb25lcyI6eyI5Ijp0cnVlLCIxMCI6dHJ1ZSwiMCI6dHJ1ZSwiMTQiOnRydWUsIjEyIjp0cnVlLCI5OSI6dHJ1ZSwiMjEiOnRydWV9fQ.tGmYBUsymQyRllRaVV9xqXhgYWXru_zKUFOQSnDHOR0",
			},
			want: 200,
		},
		{
			name: "JWTValidate - Empty token",
			args: args{
				jwt: "",
			},
			want: 403,
		},
		{
			name: "JWTValidate - Not valid token",
			args: args{
				jwt: "pepe",
			},
			want: 403,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := godotenv.Load("../.env")
			if e != nil {
				panic(e)
			}
			nextHandle := httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {})
			handlerToTest := JWTValidate(nextHandle)
			rr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/testing", nil)
			req.Header.Set("Authorization", tt.args.jwt)
			router := httprouter.New()
			router.GET("/testing", handlerToTest)
			router.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.want {
				t.Errorf("Wrong status")
			}
		})
	}
}

func TestCheckToken(t *testing.T) {
	type args struct {
		jwt    string
		secret string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "JWTValidate",
			args: args{
				jwt:    "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzZXNzaW9uVG9rZW4iOiIwMGFjNzQ1NmYxM2NiMjhhYmM2YzdjNGVkY2UwOTc4MjRjMmYxNDE2IiwiY29kY2xpIjoxMDU0NiwiY29kX3BhcnRuZXIiOjQ1Niwic29sdWNpb25lcyI6eyI5Ijp0cnVlLCIxMCI6dHJ1ZSwiMCI6dHJ1ZSwiMTQiOnRydWUsIjEyIjp0cnVlLCI5OSI6dHJ1ZSwiMjEiOnRydWV9fQ.tGmYBUsymQyRllRaVV9xqXhgYWXru_zKUFOQSnDHOR0",
				secret: "321456987",
			},
			wantErr: false,
		},
		{
			name: "JWTValidate - Empty token",
			args: args{
				jwt:    "",
				secret: "321456987",
			},
			wantErr: true,
		},
		{
			name: "JWTValidate - Empty secret",
			args: args{
				jwt:    "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzZXNzaW9uVG9rZW4iOiIwMGFjNzQ1NmYxM2NiMjhhYmM2YzdjNGVkY2UwOTc4MjRjMmYxNDE2IiwiY29kY2xpIjoxMDU0NiwiY29kX3BhcnRuZXIiOjQ1Niwic29sdWNpb25lcyI6eyI5Ijp0cnVlLCIxMCI6dHJ1ZSwiMCI6dHJ1ZSwiMTQiOnRydWUsIjEyIjp0cnVlLCI5OSI6dHJ1ZSwiMjEiOnRydWV9fQ.tGmYBUsymQyRllRaVV9xqXhgYWXru_zKUFOQSnDHOR0",
				secret: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CheckToken(tt.args.jwt, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("Wrong status")
			}
		})
	}
}

func TestGetTokenData(t *testing.T) {
	type args struct {
		jwt    string
		secret string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetTokenData",
			args: args{
				jwt:    "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzZXNzaW9uVG9rZW4iOiIwMGFjNzQ1NmYxM2NiMjhhYmM2YzdjNGVkY2UwOTc4MjRjMmYxNDE2IiwiY29kY2xpIjoxMDU0NiwiY29kX3BhcnRuZXIiOjQ1Niwic29sdWNpb25lcyI6eyI5Ijp0cnVlLCIxMCI6dHJ1ZSwiMCI6dHJ1ZSwiMTQiOnRydWUsIjEyIjp0cnVlLCI5OSI6dHJ1ZSwiMjEiOnRydWV9fQ.tGmYBUsymQyRllRaVV9xqXhgYWXru_zKUFOQSnDHOR0",
				secret: "321456987",
			},
			wantErr: false,
		},
		{
			name: "GetTokenData - Void token",
			args: args{
				jwt:    "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzZXNzaW9uVG9rZW4iOiIwMGFjNzQ1NmYxM2NiMjhhYmM2YzdjNGVkY2UwOTc4MjRjMmYxNDE2IiwiY29kY2xpIjoxMDU0NiwiY29kX3BhcnRuZXIiOjQ1Niwic29sdWNpb25lcyI6eyI5Ijp0cnVlLCIxMCI6dHJ1ZSwiMCI6dHJ1ZSwiMTQiOnRydWUsIjEyIjp0cnVlLCI5OSI6dHJ1ZSwiMjEiOnRydWV9fQ.tGmYBUsymQyRllRaVV9xqXhgYWXru_zKUFOQSnDHOR0",
				secret: "321456987",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := CheckToken(tt.args.jwt, tt.args.secret)
			if err != nil {
				panic(err)
			}
			if tt.name == "GetTokenData - Void token" {
				token.Valid = false
			}
			_, err = GetTokenData(token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Wrong status")
			}
		})
	}
}
