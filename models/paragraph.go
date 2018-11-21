package models

import (
	"fmt"
	"strings"
)

// Paragraph model for paragraphs
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

// GetParagraph Returns a single paragraph given a Paragraph.UUID
func (db *DB) GetParagraph(lang interface{}, paragraph Paragraph, include string) (*Paragraph, error) {
	p := new(Paragraph)
	includeParams := strings.Split(include, ",")
	if contains(includeParams, "compliances") {
		query := `SELECT
			paragraph_compliance.id,
			report.report_number AS report_id,
			paragraph.paragraph_number AS paragraph_id,
			p_trans.text AS primary_compliance,
			s_trans.text AS secondary_compliance,
			o_trans.text AS operational_compliance
		FROM
			paragraph_compliance
			JOIN report ON report.uuid = paragraph_compliance.report_uuid
			JOIN lkp_compliance AS p_compliance
			ON paragraph_compliance.primary_compliance = p_compliance.id
			JOIN trans_compliance AS p_trans
			ON p_trans.compliance_id = p_compliance.id
			JOIN lkp_compliance AS s_compliance
			ON paragraph_compliance.secondary_compliance = s_compliance.id
			JOIN trans_compliance AS s_trans
			ON s_trans.compliance_id = s_compliance.id
			JOIN lkp_compliance AS o_compliance
			ON paragraph_compliance.operation_compliance = o_compliance.id
			JOIN trans_compliance AS o_trans
			ON o_trans.compliance_id = o_compliance.id
			JOIN paragraph ON paragraph.uuid = paragraph_compliance.paragraph_uuid
		WHERE
			p_trans.lang_code = $1
			AND s_trans.lang_code = $1
			AND o_trans.lang_code = $1
			AND paragraph.paragraph_number = $2`
		fmt.Println(paragraph.ID)
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
	var queryBuilder strings.Builder
	queryBuilder.WriteString(paragraphQuery)
	queryBuilder.WriteString(" AND paragraph.paragraph_number = $2")
	query := queryBuilder.String()
	row := db.QueryRow(query, lang, paragraph.ID)
	err := row.Scan(&p.ID, &p.ParagraphNumber, &p.ParagraphTitle, &p.ParagraphText)
	if err != nil {
		return nil, err
	}
	return p, nil
}
