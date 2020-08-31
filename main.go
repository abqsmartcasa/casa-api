package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/abqsmartcasa/casa-api/models"

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

	paragraphs := r.PathPrefix("/paragraphs").Subrouter()
	paragraphs.Use(langCheck)
	paragraphs.Use(handleKey)
	paragraphs.HandleFunc("", env.paragraphs).Methods("GET")
	paragraphs.HandleFunc("/{key}", env.paragraph).Methods("GET")

	paragraphsRelationships := paragraphs.PathPrefix("/{key}").Subrouter()
	paragraphsRelationships.Use(langCheck)
	paragraphsRelationships.Use(handleKey)
	paragraphsRelationships.HandleFunc("/category-tags", env.categoryTagsByParagraph).Methods("GET")
	paragraphsRelationships.HandleFunc("/specific-tags", env.specificTagsByParagraph).Methods("GET")
	paragraphsRelationships.HandleFunc("/compliances", env.compliancesByParagraph).Methods("GET")

	compliances := r.PathPrefix("/compliances").Subrouter()
	compliances.Use(langCheck)
	compliances.Use(handleKey)
	compliances.HandleFunc("", env.compliances).Methods("GET")

	reports := r.PathPrefix("/reports").Subrouter()
	reports.Use(langCheck)
	reports.Use(handleKey)
	reports.HandleFunc("", env.reports).Methods("GET")
	reports.HandleFunc("/{key}", env.report).Methods("GET")
	reports.HandleFunc("/{key}/compliances", env.compliancesByReport).Methods("GET")

	categoryTags := r.PathPrefix("/category-tags").Subrouter()
	categoryTags.Use(langCheck)
	categoryTags.Use(handleKey)
	categoryTags.HandleFunc("", env.categoryTags).Methods("GET")
	categoryTags.HandleFunc("/{key}", env.categoryTag).Methods("GET")
	categoryTags.HandleFunc("/{key}/paragraphs", env.paragraphsByCategoryTag).Methods("GET")

	specificTags := r.PathPrefix("/specific-tags").Subrouter()
	specificTags.Use(langCheck)
	specificTags.Use(handleKey)
	specificTags.HandleFunc("", env.specificTags).Methods("GET")
	specificTags.HandleFunc("/{key}", env.specificTag).Methods("GET")
	specificTags.HandleFunc("/{key}/paragraphs", env.paragraphsBySpecificTag).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
