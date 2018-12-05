package models

import (
	"database/sql"

	_ "github.com/lib/pq" //postgres import
)

type Datastore interface {
	AllParagraphs(lang interface{}) ([]*Paragraph, error)
	GetParagraph(lang interface{}, paragraph Paragraph, include string) (*Paragraph, error)
	GetParagraphsBySpecificTag(lang interface{}, specificTag SpecificTag) ([]*Paragraph, error)
	GetParagraphsByCategoryTag(lang interface{}, categoryTag CategoryTag) ([]*Paragraph, error)
	AllCompliances(lang interface{}) ([]*Compliance, error)
	GetCompliance(lang interface{}, compliance Compliance) (*Compliance, error)
	GetCompliancesByParagraph(lang interface{}, paragraph Paragraph) ([]*Compliance, error)
	GetCompliancesByReport(lang interface{}, report Report) ([]*Compliance, error)
	AllReports(lang interface{}) ([]*Report, error)
	GetReport(lang interface{}, report Report) (*Report, error)
	AllCategoryTags(lang interface{}) ([]*CategoryTag, error)
	GetCategoryTag(lang interface{}, categoryTag CategoryTag) (*CategoryTag, error)
	GetCategoryTagsByParagraph(lang interface{}, paragraph Paragraph) ([]*CategoryTag, error)
	AllSpecificTags(lang interface{}) ([]*SpecificTag, error)
	GetSpecificTag(lang interface{}, specificTag SpecificTag) (*SpecificTag, error)
	GetSpecificTagsByParagraph(lang interface{}, paragraph Paragraph) ([]*SpecificTag, error)
}

type DB struct {
	*sql.DB
}

// NewDB database initializer
func NewDB(databaseURI string) (*DB, error) {
	db, err := sql.Open("postgres", databaseURI)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
