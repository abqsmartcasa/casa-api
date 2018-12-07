package models

import (
	"testing"
)

func TestAllSpecificTags(t *testing.T) {
	tests := []struct {
		description string
		rows        int
		lang        interface{}
	}{}
	for _, test := range tests {
		sts, err := TestDB.AllSpecificTags(test.lang)
		if err != nil {
			t.Errorf("")
		}
		if len(sts) != test.rows {
			t.Errorf("row count did not match")
		}
	}
}

func TestAllCategoryTags(t *testing.T) {
	tests := []struct {
		description string
		rows        int
		lang        interface{}
	}{}
	for _, test := range tests {
		cts, err := TestDB.AllCategoryTags(test.lang)
		if err != nil {
			t.Errorf("")
		}
		if len(cts) != test.rows {
			t.Errorf("row count did not match")
		}
	}
}

func TestGetSpecificTag(t *testing.T) {
	tests := []struct {
		description         string
		lang                interface{}
		specificTagID       int
		expectedSpecificTag SpecificTag
	}{}
	for _, test := range tests {
		specificTag := SpecificTag{}
		specificTag.ID = test.specificTagID
		p, err := TestDB.GetSpecificTag(test.lang, specificTag)
		if err != nil {
			t.Errorf("")
		}
		if p != &test.expectedSpecificTag {

		}
	}
}

func TestGetCategoryTag(t *testing.T) {
	tests := []struct {
		description         string
		lang                interface{}
		categoryTagID       int
		expectedCategoryTag CategoryTag
	}{}
	for _, test := range tests {
		categoryTag := CategoryTag{}
		categoryTag.ID = test.categoryTagID
		ct, err := TestDB.GetCategoryTag(test.lang, categoryTag)
		if err != nil {
			t.Errorf("")
		}
		if ct != &test.expectedCategoryTag {

		}
	}
}

func TestGetSpecificTagsByParagraph(t *testing.T) {
	tests := []struct {
		description  string
		lang         interface{}
		paragraphID  int
		expectedRows int
	}{}
	for _, test := range tests {
		paragraph := Paragraph{}
		paragraph.ID = test.paragraphID
		sts, err := TestDB.GetSpecificTagsByParagraph(test.lang, paragraph)
		if err != nil {
			t.Errorf("")
		}
		if len(sts) != test.expectedRows {

		}
	}
}

func TestGetCategoryTagsByParagraph(t *testing.T) {
	tests := []struct {
		description  string
		lang         interface{}
		paragraphID  int
		expectedRows int
	}{}
	for _, test := range tests {
		paragraph := Paragraph{}
		paragraph.ID = test.paragraphID
		cts, err := TestDB.GetCategoryTagsByParagraph(test.lang, paragraph)
		if err != nil {
			t.Errorf("")
		}
		if len(cts) != test.expectedRows {

		}
	}
}
