package models

import (
	"testing"
)

func TestAllReports(t *testing.T) {
	tests := []struct {
		description string
		rows        int
		lang        interface{}
	}{}
	for _, test := range tests {
		rpts, err := TestDB.AllReports(test.lang)
		if err != nil {
			t.Errorf("")
		}
		if len(rpts) != test.rows {
			t.Errorf("row count did not match")
		}
	}
}

func TestGetReport(t *testing.T) {
	tests := []struct {
		description    string
		lang           interface{}
		reportID       int
		expectedReport Report
	}{}
	for _, test := range tests {
		report := Report{}
		report.ID = test.reportID
		rpt, err := TestDB.GetReport(test.lang, report)
		if err != nil {
			t.Errorf("")
		}
		if rpt != &test.expectedReport {

		}
	}
}
