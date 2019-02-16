package models

import (
	"encoding/json"
	"fmt"
	"log"
)

// Compliance model for compliances
type Compliance struct {
	ReportID            int             `json:"reportId"`
	ParagraphID         int             `json:"paragraphId,omitempty"`
	PrimaryCompliance   string          `json:"primaryCompliance"`
	OperationCompliance string          `json:"operationalCompliance"`
	SecondaryCompliance string          `json:"secondaryCompliance"`
	Pages               json.RawMessage `json:"pages"`
}

// AllCompliances returns a slice with all paragraphs
func (db *DB) AllCompliances(lang interface{}) ([]*Compliance, error) {
	query := `SELECT
				report.report_number AS report_id,
				paragraph.paragraph_number AS paragraph_id,
				p_trans.text AS primary_compliance,
				s_trans.text AS secondary_compliance,
				o_trans.text AS operational_compliance,
				array_to_json(array_agg(compliance_page.page_number ORDER BY compliance_page.page_number ASC)) AS pages
			FROM
				compliance
				JOIN report ON report.uuid = compliance.report_uuid
				JOIN lkp_compliance AS p_compliance
				ON compliance.primary_compliance = p_compliance.id
				JOIN trans_compliance AS p_trans
				ON p_trans.compliance_id = p_compliance.id
				JOIN lkp_compliance AS s_compliance
				ON compliance.secondary_compliance = s_compliance.id
				JOIN trans_compliance AS s_trans
				ON s_trans.compliance_id = s_compliance.id
				JOIN lkp_compliance AS o_compliance
				ON compliance.operation_compliance = o_compliance.id
				JOIN trans_compliance AS o_trans
				ON o_trans.compliance_id = o_compliance.id
				JOIN paragraph ON paragraph.uuid = compliance.paragraph_uuid
				JOIN compliance_page ON compliance.uuid = compliance_page.compliance_uuid
			WHERE
				p_trans.lang_code = $1
				AND s_trans.lang_code = $1
				AND o_trans.lang_code = $1
			GROUP BY 1,2,3,4,5`
	rows, err := db.Query(query, lang)
	log.Println(rows)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	cs := make([]*Compliance, 0)
	for rows.Next() {
		c := new(Compliance)
		err := rows.Scan(&c.ReportID, &c.ParagraphID, &c.PrimaryCompliance, &c.SecondaryCompliance, &c.OperationCompliance, &c.Pages)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		cs = append(cs, c)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return cs, nil
}

// GetCompliancesByParagraph Returns a slice of Compliances given a Paragraph.ParagraphNumber
func (db *DB) GetCompliancesByParagraph(lang interface{}, paragraph Paragraph) ([]*Compliance, error) {
	query := `SELECT
			report.report_number AS report_id,
			paragraph.paragraph_number AS paragraph_id,
			p_trans.text AS primary_compliance,
			s_trans.text AS secondary_compliance,
			o_trans.text AS operational_compliance,
			array_to_json(array_agg(compliance_page.page_number ORDER BY compliance_page.page_number ASC)) AS pages
		FROM
			compliance
			JOIN report ON report.uuid = compliance.report_uuid
			JOIN lkp_compliance AS p_compliance
			ON compliance.primary_compliance = p_compliance.id
			JOIN trans_compliance AS p_trans
			ON p_trans.compliance_id = p_compliance.id
			JOIN lkp_compliance AS s_compliance
			ON compliance.secondary_compliance = s_compliance.id
			JOIN trans_compliance AS s_trans
			ON s_trans.compliance_id = s_compliance.id
			JOIN lkp_compliance AS o_compliance
			ON compliance.operation_compliance = o_compliance.id
			JOIN trans_compliance AS o_trans
			ON o_trans.compliance_id = o_compliance.id
			JOIN paragraph ON paragraph.uuid = compliance.paragraph_uuid
			JOIN compliance_page ON compliance.uuid = compliance_page.compliance_uuid
		WHERE
			p_trans.lang_code = $1
			AND s_trans.lang_code = $1
			AND o_trans.lang_code = $1
			AND paragraph.paragraph_number = $2
		GROUP BY 1,2,3,4,5`
	rows, err := db.Query(query, lang, paragraph.ID)
	if err != nil {
		return nil, err
	}
	cs := make([]*Compliance, 0)
	defer rows.Close()
	for rows.Next() {
		c := new(Compliance)
		err := rows.Scan(&c.ReportID, &c.ParagraphID, &c.PrimaryCompliance, &c.SecondaryCompliance, &c.OperationCompliance, &c.Pages)
		if err != nil {
			return nil, err
		}
		cs = append(cs, c)
	}
	return cs, nil
}

// GetCompliancesByReport Returns a slice of Compliances given a Report.ID
func (db *DB) GetCompliancesByReport(lang interface{}, report Report) ([]*Compliance, error) {
	query := `SELECT
			report.report_number AS report_id,
			paragraph.paragraph_number AS paragraph_id,
			p_trans.text AS primary_compliance,
			s_trans.text AS secondary_compliance,
			o_trans.text AS operational_compliance,
			array_to_json(array_agg(compliance_page.page_number ORDER BY compliance_page.page_number ASC)) AS pages
		FROM
			compliance
			JOIN report ON report.uuid = compliance.report_uuid
			JOIN lkp_compliance AS p_compliance
			ON compliance.primary_compliance = p_compliance.id
			JOIN trans_compliance AS p_trans
			ON p_trans.compliance_id = p_compliance.id
			JOIN lkp_compliance AS s_compliance
			ON compliance.secondary_compliance = s_compliance.id
			JOIN trans_compliance AS s_trans
			ON s_trans.compliance_id = s_compliance.id
			JOIN lkp_compliance AS o_compliance
			ON compliance.operation_compliance = o_compliance.id
			JOIN trans_compliance AS o_trans
			ON o_trans.compliance_id = o_compliance.id
			JOIN paragraph ON paragraph.uuid = compliance.paragraph_uuid
			JOIN compliance_page ON compliance.uuid = compliance_page.compliance_uuid
		WHERE
			p_trans.lang_code = $1
			AND s_trans.lang_code = $1
			AND o_trans.lang_code = $1
			AND report.report_number = $2
		GROUP BY 1,2,3,4,5`
	rows, err := db.Query(query, lang, report.ID)
	if err != nil {
		return nil, err
	}
	cs := make([]*Compliance, 0)
	defer rows.Close()
	for rows.Next() {
		c := new(Compliance)
		err := rows.Scan(&c.ReportID, &c.ParagraphID, &c.PrimaryCompliance, &c.SecondaryCompliance, &c.OperationCompliance, &c.Pages)
		if err != nil {
			return nil, err
		}
		cs = append(cs, c)
	}
	return cs, nil
}
