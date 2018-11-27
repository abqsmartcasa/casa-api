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

func handler(w http.ResponseWriter, r *http.Request) {
	return
}

func TestIncludeCheck(t *testing.T) {
	icm := includeCheck{validParams: []string{"compliances"}}
	var tests = []struct {
		URL          string
		expectedCode int
		description  string
	}{
		{"/paragraphs/13?include=compliances", 200, "key with compliances include"},
		{"/paragraphs/13", 200, "key alone"},
		{"/paragraphs?include=compliances", 400, "without key with include compliances"},
		{"/paragraphs/13?include=somethingelse", 400, "key with single invalid include"},
		{"/paragraphs/13?include=compliances,reports", 400, "key with single invalid include in list"},
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	r := router.PathPrefix("/paragraphs").Subrouter()
	r.Use(icm.Middleware)
	r.HandleFunc("", handler)
	r.HandleFunc("/{key}", handler)

	for _, tc := range tests {
		req, err := http.NewRequest("GET", tc.URL, nil)
		if err != nil {
			t.Fatal(err)
		}

		r.ServeHTTP(rr, req)

		if rr.Code != tc.expectedCode {
			t.Errorf("%s: got %v want %v",
				tc.description, rr.Code, tc.expectedCode)
		}
	}
}

func TestIncludeCheckNoParams(t *testing.T) {
	icm := includeCheck{validParams: []string{}}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	r := router.PathPrefix("/paragraphs").Subrouter()
	r.Use(icm.Middleware)
	r.HandleFunc("", handler)
	r.HandleFunc("/{key}", handler)
	req, err := http.NewRequest("GET", "/paragraphs/13?include=compliances", nil)
	if err != nil {
		t.Fatal(err)
	}

	r.ServeHTTP(rr, req)

	if rr.Code != 400 {
		t.Errorf(" got %v want %v",
			rr.Code, 400)
	}
}
