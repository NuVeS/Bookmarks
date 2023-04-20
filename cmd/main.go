package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"Bookmarks/cmd/router"

	"github.com/jackc/pgx/v5"
)

func main() {
	databaseUrl := "postgres://admin:123456@localhost:5432/postgres"
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	r := router.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}
