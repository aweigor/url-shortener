package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/configs"
	"url-shortener/internal/auth"
	"url-shortener/internal/user"
	"url-shortener/pkg/db"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func bootstrap() (*auth.AuthHandler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))
	if err != nil {
		return nil, nil, err
	}
	userRepo := user.NewUserRepository(&db.Db{
		DB: gormDb,
	})
	handler := auth.AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		AuthService: auth.NewAuthService(userRepo),
	}
	return &handler, mock, nil
}

func TestLoginSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	rows := sqlmock.NewRows([]string{"email", "password"}).AddRow("mail.mail@mail", "$2a$10$BC6XX/I4TlmhixGq/zPJnO60uN5fy8GZp7AVniXXX.iO8NPs.A6P2")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	if err != nil {
		t.Fatal("Failed init testing")
		return
	}
	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "mail.mail@mail",
		Password: "123", // $2a$10$BC6XX/I4TlmhixGq/zPJnO60uN5fy8GZp7AVniXXX.iO8NPs.A6P2
	})
	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	handler.Login()(w, req)
}
