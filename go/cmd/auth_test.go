package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"url-shortener/internal/auth"
	"url-shortener/internal/user"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initData(db *gorm.DB) {
	// todo: get password hash from existing record
	db.Create(&user.User{
		Email:    "test@mail.test",
		Password: "12345",
		Name:     "john",
	})
}

func cleanData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?", "test@mail.test").
		Delete(&user.User{})
}

func TestLoginSuccess(t *testing.T) {

	db := initDB()
	initData(db)

	ts := httptest.NewServer(NewApp())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a2@a.ru",
		Password: "1",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Expected %d got %d", 200, res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var resData auth.LoginResponse
	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal()
	}
	if resData.Token == "" {
		t.Fatal("Token not defined")
	}
	cleanData(db)
}

func TestLoginFail(t *testing.T) {
	db := initDB()
	initData(db)

	ts := httptest.NewServer(NewApp())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a2@a.ru",
		Password: "1",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 401 {
		t.Fatalf("Expected %d got %d", 401, res.StatusCode)
	}
	cleanData(db)
}
