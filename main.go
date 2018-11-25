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
	paragraphIncludes := includeCheck{validParams: []string{"compliances", "tags"}}

	paragraphs := r.PathPrefix("/paragraphs").Subrouter()
	paragraphs.Use(langCheck)
	paragraphs.Use(paragraphIncludes.Middleware)
	paragraphs.HandleFunc("", env.paragraphs).Methods("GET")
	paragraphs.HandleFunc("/{key}", env.paragraph).Methods("GET")

	paragraphsTags := r.PathPrefix("/paragraphs/{key}").Subrouter()
	paragraphTagIncludes := includeCheck{validParams: []string{}}
	paragraphsTags.Use(langCheck)
	paragraphsTags.Use(paragraphTagIncludes.Middleware)
	paragraphsTags.HandleFunc("/categorytags", env.categoryTagsByParagraph).Methods("GET")
	paragraphsTags.HandleFunc("/specifictags", env.specificTagsByParagraph).Methods("GET")

	complianceIncludes := includeCheck{validParams: []string{}}
	compliances := r.PathPrefix("/compliances").Subrouter()
	compliances.Use(langCheck)
	compliances.Use(complianceIncludes.Middleware)
	compliances.HandleFunc("", env.compliances).Methods("GET")
	compliances.HandleFunc("/{key}", env.compliance).Methods("GET")

	reports := r.PathPrefix("/reports").Subrouter()
	reportsIncludes := includeCheck{validParams: []string{}}
	reports.Use(langCheck)
	reports.Use(reportsIncludes.Middleware)
	reports.HandleFunc("", env.reports).Methods("GET")
	reports.HandleFunc("/{key}", env.report).Methods("GET")

	categoryTags := r.PathPrefix("/categorytags").Subrouter()
	categoryTagsIncludes := includeCheck{validParams: []string{"specifictags"}}
	categoryTags.Use(langCheck)
	categoryTags.Use(categoryTagsIncludes.Middleware)
	categoryTags.HandleFunc("", env.categoryTags).Methods("GET")
	categoryTags.HandleFunc("/{key}", env.categoryTag).Methods("GET")

	categoryTags.HandleFunc("/{key}/paragraphs", env.paragraphsByCategoryTag).Methods("GET")

	specificTags := r.PathPrefix("/specifictags").Subrouter()
	specificTagsIncludes := includeCheck{validParams: []string{}}
	specificTags.Use(langCheck)
	specificTags.Use(specificTagsIncludes.Middleware)
	specificTags.HandleFunc("", env.specificTags).Methods("GET")
	specificTags.HandleFunc("/{key}", env.specificTag).Methods("GET")

	specificTagsParagraphs := r.PathPrefix("/specifictags/{key}").Subrouter()
	specificTagsParagraphsIncludes := includeCheck{validParams: []string{}}
	specificTags.Use(langCheck)
	specificTags.Use(specificTagsParagraphsIncludes.Middleware)
	specificTagsParagraphs.HandleFunc("/paragraphs", env.paragraphsBySpecificTag).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
