package router

import (
	"fmt"
	"net/http"

	"Bookmarks/cmd/auth"
	"Bookmarks/cmd/logger"
	"Bookmarks/cmd/storage"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(pool *pgxpool.Pool) *chi.Mux {
	router := chi.NewRouter()
	storage := &storage.Storage{pool}

	auth := auth.NewAuthHandler(storage)

	router.Post("//register", auth.Register)

	router.Use(logger.Logger)
	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
