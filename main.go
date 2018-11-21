package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/apdforward/apdf_api/models"

	"github.com/gorilla/mux"
)

func main() {
	userPassword := url.UserPassword(os.Getenv("USER"), os.Getenv("PASSWORD"))
	URI := new(url.URL)
	URI.Scheme = "postgres"
	URI.User = userPassword
	URI.Host = os.Getenv("HOST")
	URI.Path = os.Getenv("DBNAME")
	URI.RawQuery = "sslmode=disable"
	db, err := models.NewDB(URI.String())
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db: db}
	r := mux.NewRouter()
	paragraphIncludes := includeCheck{validParams: []string{"compliances"}}
	paragraphs := r.PathPrefix("/paragraphs").Subrouter()
	paragraphs.Use(langCheck)
	paragraphs.Use(paragraphIncludes.Middleware)
	paragraphs.HandleFunc("", env.paragraphs).Methods("GET")
	paragraphs.HandleFunc("/{key}", env.paragraph).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
