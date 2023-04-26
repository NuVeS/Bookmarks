package auth

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"Bookmarks/cmd/storage"
)

type RegisterRequest struct {
	Login      string `json:"login"`
	Password   string `json:"password"`
	RepeatPass string `json:"repeat_pass"`
}

type Auth struct {
	storage storage.API
}

func NewAuthHandler(storage storage.API) (auth *Auth) {
	auth = &Auth{storage: storage}
	return
}

func (auth *Auth) Register(w http.ResponseWriter, r *http.Request) {
	var body RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sendError(w)
		return
	}

	if body.Password != body.RepeatPass {
		w.WriteHeader(http.StatusBadRequest)
		sendError(w)
		return
	}

	md5 := md5.Sum([]byte(body.Password))
	sum := hex.EncodeToString(md5[:])
	err = auth.storage.AddUser(context.Background(), body.Login, sum)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		sendError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func sendError(writer http.ResponseWriter) {
	json, err := json.Marshal("Failed")
	if err == nil {
		writer.Write(json)
	}
}
