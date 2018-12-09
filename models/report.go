package models

// Report model for Independent monitoring reports
type Report struct {
	UUID        string `json:"-"`
	ID          int    `json:"id"`
	ReportName  string `json:"report_name"`
	ReportTitle string `json:"report_title"`
	PublishDate string `json:"publish_date"`
	PeriodBegin string `json:"period_begin"`
	PeriodEnd   string `json:"period_end"`
}

var reportQuery = `SELECT
			report."report_number" as id,
			'IMR-' || report."report_number" as report_name,
			trans_report_title."text" as report_title,
			to_char(report."publish_date", 'YYYY-MM-DD'),
			to_char(report."period_begin", 'YYYY-MM-DD'),
			to_char(report."period_end", 'YYYY-MM-DD')
		FROM
			report
			JOIN trans_report_title 
			ON report.uuid = trans_report_title.report_uuid
		WHERE 
			trans_report_title.lang_code = $1`

// AllReports returns a slice with all paragraphs
func (db *DB) AllReports(lang interface{}) ([]*Report, error) {
	query := `SELECT
			report."report_number" as id,
			'IMR-' || report."report_number" as report_name,
			trans_report_title."text" as report_title,
			to_char(report."publish_date", 'YYYY-MM-DD'),
			to_char(report."period_begin", 'YYYY-MM-DD'),
			to_char(report."period_end", 'YYYY-MM-DD')
		FROM
			report
			JOIN trans_report_title 
			ON report.uuid = trans_report_title.report_uuid
		WHERE 
			trans_report_title.lang_code = $1`
	rows, err := db.Query(query, lang)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rpts := make([]*Report, 0)
	for rows.Next() {
		rpt := new(Report)
		err := rows.Scan(&rpt.ID, &rpt.ReportName, &rpt.ReportTitle, &rpt.PublishDate, &rpt.PeriodBegin, &rpt.PeriodEnd)
		if err != nil {
			return nil, err
		}
		rpts = append(rpts, rpt)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return rpts, nil
}

// GetReport Returns a single paragraph given a Report.ID (report_number in DB)
func (db *DB) GetReport(lang interface{}, report Report) (*Report, error) {
	rpt := new(Report)
	query := `SELECT
			report."report_number" as id,
			'IMR-' || report."report_number" as report_name,
			trans_report_title."text" as report_title,
			to_char(report."publish_date", 'YYYY-MM-DD'),
			to_char(report."period_begin", 'YYYY-MM-DD'),
			to_char(report."period_end", 'YYYY-MM-DD')
		FROM
			report
			JOIN trans_report_title 
			ON report.uuid = trans_report_title.report_uuid
		WHERE 
			trans_report_title.lang_code = $1
			AND report.report_number = $2`
	row := db.QueryRow(query, lang, report.ID)
	err := row.Scan(&rpt.ID, &rpt.ReportName, &rpt.ReportTitle, &rpt.PublishDate, &rpt.PeriodBegin, &rpt.PeriodEnd)
	if err != nil {
		return nil, err
	}
	return rpt, nil
}
