package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apdforward/apdf_api/models"
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

func TestParagraphs(t *testing.T) {
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/paragraphs", nil)
	if err != nil {
		fmt.Printf("error with request: %v", err)
	}
	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.paragraphs).ServeHTTP(rec, req)
	expected := "{\"data\":[{\"id\":42,\"paragraph_number\":42,\"paragraph_title\":\"test\",\"paragraph_text\":\"test\"}]}"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}
