package models

// Compliance model for compliances
type Compliance struct {
	UUID                string `json:"-"`
	ID                  int    `json:"id"`
	ReportID            string `json:"report_id"`
	ParagraphID         string `json:"paragraph_id"`
	PrimaryCompliance   string `json:"primary_compliance"`
	OperationCompliance string `json:"operational_compliance"`
	SecondaryCompliance string `json:"secondary_compliance"`
}
