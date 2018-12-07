package models

import (
	"testing"
)

func TestAllCompliances(t *testing.T) {
	tests := []struct {
		description string
		rows        int
		lang        interface{}
	}{}
	for _, test := range tests {
		cs, err := TestDB.AllCompliances(test.lang)
		if err != nil {
			t.Errorf("")
		}
		if len(cs) != test.rows {
			t.Errorf("row count did not match")
		}
	}
}

func TestGetCompliance(t *testing.T) {
	tests := []struct {
		description        string
		lang               interface{}
		complianceID       int
		expectedCompliance Compliance
		includes           string
	}{}
	for _, test := range tests {
		compliance := Compliance{}
		compliance.ID = test.complianceID
		p, err := TestDB.GetCompliance(test.lang, compliance)
		if err != nil {
			t.Errorf("")
		}
		if p != &test.expectedCompliance {

		}
	}
}

func TestGetCompliancesByParagraph(t *testing.T) {
	tests := []struct {
		description  string
		lang         interface{}
		paragraphID  int
		expectedRows int
	}{}
	for _, test := range tests {
		paragraph := Paragraph{}
		paragraph.ID = test.paragraphID
		cs, err := TestDB.GetCompliancesByParagraph(test.lang, paragraph)
		if err != nil {
			t.Errorf("")
		}
		if len(cs) != test.expectedRows {

		}
	}
}

func TestGetCompliancesByReport(t *testing.T) {
	tests := []struct {
		description  string
		lang         interface{}
		reportID     int
		expectedRows int
	}{}
	for _, test := range tests {
		report := Report{}
		report.ID = test.reportID
		cs, err := TestDB.GetCompliancesByReport(test.lang, report)
		if err != nil {
			t.Errorf("")
		}
		if len(cs) != test.expectedRows {

		}
	}
}
