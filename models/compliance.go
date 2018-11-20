package models

// Compliance model for compliances
type Compliance struct {
	UUID                *string `json:"id"`
	ReportUUID          *string `json:"report_id"`
	ReportName          *string `json:"report_name"`
	ParagraphUUID       *string `json:"paragraph_id"`
	PrimaryCompliance   *string `json:"primary_compliance"`
	OperationCompliance *string `json:"operational_compliance"`
	SecondaryCompliance *string `json:"secondary_compliance"`
}
