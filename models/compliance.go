package models

import "strings"

// Compliance model for compliances
type Compliance struct {
	UUID                string `json:"-"`
	ID                  int    `json:"id"`
	ReportID            int    `json:"report_id"`
	ParagraphID         int    `json:"paragraph_id"`
	PrimaryCompliance   string `json:"primary_compliance"`
	OperationCompliance string `json:"operational_compliance"`
	SecondaryCompliance string `json:"secondary_compliance"`
}

var complianceQuery = `SELECT
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
			AND o_trans.lang_code = $1`

// AllCompliances returns a slice with all paragraphs
func (db *DB) AllCompliances(lang interface{}) ([]*Compliance, error) {
	rows, err := db.Query(complianceQuery, lang)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cs := make([]*Compliance, 0)
	for rows.Next() {
		c := new(Compliance)
		err := rows.Scan(&c.ID, &c.ReportID, &c.ParagraphID, &c.PrimaryCompliance, &c.SecondaryCompliance, &c.OperationCompliance)
		if err != nil {
			return nil, err
		}
		cs = append(cs, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return cs, nil
}

// GetCompliance Returns a single paragraph given a Paragraph.UUID
func (db *DB) GetCompliance(lang interface{}, compliance Compliance) (*Compliance, error) {
	c := new(Compliance)
	var cQueryBuilder strings.Builder
	cQueryBuilder.WriteString(complianceQuery)
	cQueryBuilder.WriteString(" AND paragraph_compliance.id = $2")
	query := cQueryBuilder.String()
	row := db.QueryRow(query, lang, compliance.ID)
	err := row.Scan(&c.ID, &c.ReportID, &c.ParagraphID, &c.PrimaryCompliance, &c.SecondaryCompliance, &c.OperationCompliance)
	if err != nil {
		return nil, err
	}
	return c, nil
}
