package models

import (
	"testing"
)

func TestAllParagraphs(t *testing.T) {
	tests := []struct {
		description string
		rows        int
		lang        interface{}
	}{
		{"", 1, "en"},
		{"", 0, 1},
		{"", 1, "es"},
	}
	for _, test := range tests {
		ps, err := TestDB.AllParagraphs(test.lang)
		if err != nil {
			t.Errorf("")
		}
		if len(ps) != test.rows {
			t.Errorf("row count did not match")
		}
	}
}

func TestGetParagraph(t *testing.T) {
	tests := []struct {
		description       string
		lang              interface{}
		paragraphID       int
		expectedParagraph Paragraph
	}{}
	for _, test := range tests {
		paragraph := Paragraph{}
		paragraph.ID = test.paragraphID
		p, err := TestDB.GetParagraph(test.lang, paragraph)
		if err != nil {
			t.Errorf("")
		}
		if p != &test.expectedParagraph {

		}
	}
}

func TestGetParagraphsBySpecificTag(t *testing.T) {
	tests := []struct {
		description   string
		lang          interface{}
		specificTagID int
		expectedRows  int
	}{}
	for _, test := range tests {
		specificTag := SpecificTag{}
		specificTag.ID = test.specificTagID
		ps, err := TestDB.GetParagraphsBySpecificTag(test.lang, specificTag)
		if err != nil {
			t.Errorf("")
		}
		if len(ps) != test.expectedRows {

		}
	}
}

func TestGetParagraphsByCategoryTag(t *testing.T) {
	tests := []struct {
		description   string
		lang          interface{}
		categoryTagID int
		expectedRows  int
	}{}
	for _, test := range tests {
		categoryTag := CategoryTag{}
		categoryTag.ID = test.categoryTagID
		ps, err := TestDB.GetParagraphsByCategoryTag(test.lang, categoryTag)
		if err != nil {
			t.Errorf("")
		}
		if len(ps) != test.expectedRows {

		}
	}
}
