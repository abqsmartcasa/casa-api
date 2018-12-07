package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
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

func TestHandleKey(t *testing.T) {
	var tests = []struct {
		key  string
		code int
	}{
		{"/14", 200},
		{"/test", 400},
	}

	for _, test := range tests {
		testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
		rr := httptest.NewRecorder()
		handler := handleKey(testHandler)
		router := mux.NewRouter()
		router.Handle("/{key}", handler)
		req := httptest.NewRequest("GET", test.key, nil)
		router.ServeHTTP(rr, req)
		if test.code != rr.Code {
			t.Errorf("\nhandler returned wrong status code: got %v want %v",
				rr.Code, test.code)
		}
	}
}
