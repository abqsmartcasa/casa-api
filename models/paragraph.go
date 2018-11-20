package models

import "strings"

// Paragraph model for paragraphs
type Paragraph struct {
	UUID            string        `json:"-"`
	ID              int           `json:"id"`
	ParagraphNumber int32         `json:"paragraph_number"`
	ParagraphTitle  string        `json:"paragraph_title"`
	ParagraphText   string        `json:"paragraph_text"`
	CategoryTag     *CategoryTag  `json:"tags,omitempty"`
	Compliances     []*Compliance `json:"compliances,omitempty"`
}

func paragraphQuery() strings.Builder {
	var queryBuilder strings.Builder
	queryBuilder.WriteString(`SELECT
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
						"trans_paragraph_text"."lang_code" = $1`)
	return queryBuilder
}

// AllParagraphs returns a slice with all paragraphs
func (db *DB) AllParagraphs(lang interface{}) ([]*Paragraph, error) {
	query := paragraphQuery()
	rows, err := db.Query(query.String(), lang)
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
