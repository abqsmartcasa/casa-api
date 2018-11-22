package models

import (
	"fmt"
	"strings"
)

// Paragraph model for CASA paragraphs
type Paragraph struct {
	UUID            string        `json:"-"`
	ID              int           `json:"id"`
	ParagraphNumber int           `json:"paragraph_number"`
	ParagraphTitle  string        `json:"paragraph_title"`
	ParagraphText   string        `json:"paragraph_text"`
	CategoryTag     *CategoryTag  `json:"tags,omitempty"`
	Compliances     []*Compliance `json:"compliances,omitempty"`
}

var paragraphQuery = `SELECT
paragraph."paragraph_number" AS "id",
paragraph."paragraph_number" AS "paragraph_number",
"trans_paragraph_description"."text" AS "paragraph_title",
"trans_paragraph_text"."text" AS "paragraph_text"
FROM
paragraph
LEFT JOIN "trans_paragraph_description"
ON "trans_paragraph_description"."paragraph_uuid" = paragraph.uuid 
LEFT JOIN "trans_paragraph_text"
ON "trans_paragraph_text"."paragraph_uuid" = paragraph.uuid 
WHERE
"trans_paragraph_text"."lang_code" = $1`

// AllParagraphs returns a slice with all paragraphs
func (db *DB) AllParagraphs(lang interface{}) ([]*Paragraph, error) {
	rows, err := db.Query(paragraphQuery, lang)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ps := make([]*Paragraph, 0)
	for rows.Next() {
		p := new(Paragraph)
		err := rows.Scan(&p.ID, &p.ParagraphNumber, &p.ParagraphTitle, &p.ParagraphText)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ps, nil
}

// GetParagraph Returns a single paragraph given a Paragraph.ID
func (db *DB) GetParagraph(lang interface{}, paragraph Paragraph, include string) (*Paragraph, error) {
	p := new(Paragraph)
	includeParams := strings.Split(include, ",")
	if contains(includeParams, "compliances") {
		var cQueryBuilder strings.Builder
		cQueryBuilder.WriteString(complianceQuery)
		cQueryBuilder.WriteString(" AND paragraph.paragraph_number = $2")
		query := cQueryBuilder.String()
		rows, err := db.Query(query, lang, paragraph.ID)
		cs := make([]*Compliance, 0)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			c := new(Compliance)
			err := rows.Scan(&c.ID, &c.ReportID, &c.ParagraphID, &c.PrimaryCompliance, &c.SecondaryCompliance, &c.OperationCompliance)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			cs = append(cs, c)
		}
		if err = rows.Err(); err != nil {
			return nil, err
		}
		p.Compliances = cs
	}
	var pQueryBuilder strings.Builder
	pQueryBuilder.WriteString(paragraphQuery)
	pQueryBuilder.WriteString(" AND paragraph.paragraph_number = $2")
	query := pQueryBuilder.String()
	row := db.QueryRow(query, lang, paragraph.ID)
	err := row.Scan(&p.ID, &p.ParagraphNumber, &p.ParagraphTitle, &p.ParagraphText)
	if err != nil {
		return nil, err
	}
	return p, nil
}
