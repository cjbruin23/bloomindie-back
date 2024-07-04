package main

import (
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

type usersResource struct{}

func (rs usersResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.List)
	r.Post("/", rs.Create)

	return r
}

func (rs usersResource) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("users list of stuff.."))
}

func (rs usersResource) Create(w http.ResponseWriter, r *http.Request) {
	io.Copy(os.Stdout, r.Body)
	w.Write([]byte("Hello World!"))
}
