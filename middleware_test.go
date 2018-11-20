package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/context"
)

func TestLangCheck(t *testing.T) {
	var tests = []struct {
		acceptLanguage string
		outLang        string
	}{
		{"en-US", "en"},
		{"es", "es"},
		{"tr", "en"},
	}
	for _, test := range tests {
		testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			val := context.Get(r, "lang")
			if val != test.outLang {
				t.Error(fmt.Sprintf("language header not %s", test.outLang))
			}
		})
		handler := langCheck(testHandler)
		req := httptest.NewRequest("GET", "http://testing", nil)
		req.Header.Add("Accept-Language", test.acceptLanguage)
		handler.ServeHTTP(httptest.NewRecorder(), req)
	}
}
