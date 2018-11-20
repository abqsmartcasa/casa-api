package main

import (
	"net/http"

	"github.com/gorilla/context"
	"golang.org/x/text/language"
)

type middleware func(http.HandlerFunc) http.HandlerFunc

var matcher = language.NewMatcher([]language.Tag{
	language.English,
	language.AmericanEnglish,
	language.Spanish,
})

func langCheck(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang, _ := r.Cookie("lang")
		accept := r.Header.Get("Accept-Language")
		tag, _ := language.MatchStrings(matcher, lang.String(), accept)
		if tag.String() == "en-US" {
			context.Set(r, "lang", "en")
		} else {
			context.Set(r, "lang", tag.String())
		}
		f(w, r)
	}
}
