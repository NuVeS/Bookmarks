package auth

import (
	"Bookmarks/cmd/storage"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

type routeFunc = func(w http.ResponseWriter, r *http.Request)

func init() {
	databaseUrl := "postgres://admin:123456@localhost:5432/bookmarks"
	pool, _ = pgxpool.New(context.Background(), databaseUrl)
}

func TestRegister(t *testing.T) {
	storage := &storage.Storage{pool}

	auth := NewAuthHandler(storage)

	rr := makeTest(t,
		"POST",
		"/register",
		"",
		"{\"login\": \"root\",\"password\": \"123\", \"repeat_pass\": \"123\"}",
		auth.Register,
	)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func makeTest(t *testing.T, method string, path string, token string, body string, route routeFunc) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	if token != "" {
		req.Header.Add("token", token)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route(w, r)
	})

	handler.ServeHTTP(rr, req)
	return rr
}
