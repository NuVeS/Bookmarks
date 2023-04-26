package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"Bookmarks/cmd/router"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	databaseUrl := "postgres://admin:123456@localhost:5432/bookmarks"
	conn, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	r := router.NewRouter(conn)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}
