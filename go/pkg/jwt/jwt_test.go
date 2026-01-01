package jwt

import (
	"testing"
)

func TestJWTCreate(t *testing.T) {
	const email = "mail@test.com"
	jwtService := NewJWT("token")
	token, err := jwtService.Create(JWTData{
		Email: "mail@test.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Token is invalid")
	}
	if data.Email != email {
		t.Fatalf("Email %s not equal %s", data.Email, email)
	}
}
