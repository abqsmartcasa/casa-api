package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apdforward/apdf_api/models"
	"github.com/gorilla/mux"
)

type mockDB struct{}

func (mdb *mockDB) AllParagraphs(lang interface{}) ([]*models.Paragraph, error) {
	ps := make([]*models.Paragraph, 0)
	ps = append(ps, &models.Paragraph{
		ID:              42,
		ParagraphNumber: 42,
		ParagraphTitle:  "test",
		ParagraphText:   "test",
	})
	return ps, nil
}

func (mdb *mockDB) GetParagraph(lang interface{}, paragraph models.Paragraph, include string) (*models.Paragraph, error) {
	p := &models.Paragraph{
		ID:              13,
		ParagraphNumber: 13,
		ParagraphTitle:  "test",
		ParagraphText:   "test",
	}
	return p, nil
}

func TestParagraphs(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/paragraphs", nil)
	if err != nil {
		t.Fatal(err)
	}
	router := mux.NewRouter()
	env := Env{db: &mockDB{}}
	router.HandleFunc("/paragraphs", env.paragraphs)
	router.ServeHTTP(rr, req)
	expected := "{\"data\":[{\"id\":42,\"paragraph_number\":42,\"paragraph_title\":\"test\",\"paragraph_text\":\"test\"}]}"
	if expected != rr.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rr.Body.String())
	}
}

func TestParagraph(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/paragraphs/13", nil)
	if err != nil {
		t.Fatal(err)
	}
	router := mux.NewRouter()
	env := Env{db: &mockDB{}}
	router.HandleFunc("/paragraphs/{key}", env.paragraph)
	router.ServeHTTP(rr, req)
	expected := "{\"data\":{\"id\":13,\"paragraph_number\":13,\"paragraph_title\":\"test\",\"paragraph_text\":\"test\"}}"
	if expected != rr.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rr.Body.String())
	}

}
