package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/gorilla/context"
	"golang.org/x/text/language"
)

var matcher = language.NewMatcher([]language.Tag{
	language.English,
	language.AmericanEnglish,
	language.Spanish,
})

func langCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang, _ := r.Cookie("lang")
		accept := r.Header.Get("Accept-Language")
		tag, _ := language.MatchStrings(matcher, lang.String(), accept)
		if tag.String() == "en-US" {
			context.Set(r, "lang", "en")
		} else {
			context.Set(r, "lang", tag.String())
		}
		next.ServeHTTP(w, r)
	})
}

func handleKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if len(vars) > 0 {
			key, err := strconv.Atoi(vars["key"])
			if err != nil {
				http.Error(w, http.StatusText(400), 400)
				return
			}
			context.Set(r, "key", key)
		}
		next.ServeHTTP(w, r)
	})
}
