package models

import (
	"database/sql"

	_ "github.com/lib/pq" //postgres import
)

type Datastore interface {
	AllParagraphs(lang interface{}) ([]*Paragraph, error)
	GetParagraph(lang interface{}, paragraph Paragraph, include string) (*Paragraph, error)
	AllCompliances(lang interface{}) ([]*Compliance, error)
	GetCompliance(lang interface{}, compliance Compliance) (*Compliance, error)
	AllReports(lang interface{}) ([]*Report, error)
	GetReport(lang interface{}, report Report) (*Report, error)
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
