package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apdforward/apdf_api/models"
	"github.com/gorilla/mux"
)

type mockDB struct{}

func (mdb *mockDB) AllParagraphs(lang interface{}) ([]*models.Paragraph, error) {
	ps := make([]*models.Paragraph, 0)
	ps = append(ps, &models.Paragraph{
		ID:              42,
		ParagraphNumber: 42,
		ParagraphTitle:  "test",
		ParagraphText:   "test",
	})
	return ps, nil
}

func (mdb *mockDB) GetParagraph(lang interface{}, paragraph models.Paragraph, include string) (*models.Paragraph, error) {
	p := &models.Paragraph{
		ID:              13,
		ParagraphNumber: 13,
		ParagraphTitle:  "test",
		ParagraphText:   "test",
	}
	return p, nil
}

func (mdb *mockDB) AllCompliances(lang interface{}) ([]*models.Compliance, error) {
	cs := make([]*models.Compliance, 0)
	cs = append(cs, &models.Compliance{
		ID:                  1,
		ReportID:            2,
		ParagraphID:         3,
		PrimaryCompliance:   "In Compliance",
		SecondaryCompliance: "Not In Compliance",
		OperationCompliance: "Not In Compliance",
	})
	return cs, nil
}

func (mdb *mockDB) GetCompliance(lang interface{}, compliance models.Compliance) (*models.Compliance, error) {
	c := &models.Compliance{
		ID:                  13,
		ReportID:            2,
		ParagraphID:         3,
		PrimaryCompliance:   "In Compliance",
		SecondaryCompliance: "Not In Compliance",
		OperationCompliance: "Not In Compliance",
	}
	return c, nil
}

func (mdb *mockDB) AllReports(lang interface{}) ([]*models.Report, error) {
	rpts := make([]*models.Report, 0)
	rpts = append(rpts, &models.Report{
		ID:          1,
		ReportName:  "IMR-1",
		ReportTitle: "Monitor's First Report",
		PublishDate: "2015-11-23",
		PeriodBegin: "2015-02-01",
		PeriodEnd:   "2015-05-31",
	})
	return rpts, nil
}

func (mdb *mockDB) GetReport(lang interface{}, report models.Report) (*models.Report, error) {
	c := &models.Report{
		ID:          1,
		ReportName:  "IMR-1",
		ReportTitle: "Monitor's First Report",
		PublishDate: "2015-11-23",
		PeriodBegin: "2015-02-01",
		PeriodEnd:   "2015-05-31",
	}
	return c, nil
}

func TestHandlers(t *testing.T) {
	router := mux.NewRouter()
	env := Env{db: &mockDB{}}
	router.HandleFunc("/paragraphs", env.paragraphs)
	router.HandleFunc("/paragraphs/{key}", env.paragraph)
	router.HandleFunc("/compliances", env.compliances)
	router.HandleFunc("/compliances/{key}", env.compliance)
	router.HandleFunc("/reports", env.reports)
	router.HandleFunc("/reports/{key}", env.report)

	tests := []struct {
		description string
		URL         string
		expected    string
	}{
		{"all paragraphs", "/paragraphs", "{\"data\":[{\"id\":42,\"paragraph_number\":42,\"paragraph_title\":\"test\",\"paragraph_text\":\"test\"}]}"},
		{"paragraph by key", "/paragraphs/13", "{\"data\":{\"id\":13,\"paragraph_number\":13,\"paragraph_title\":\"test\",\"paragraph_text\":\"test\"}}"},
		{"all compliances", "/compliances", "{\"data\":[{\"id\":1,\"report_id\":2,\"paragraph_id\":3,\"primary_compliance\":\"In Compliance\",\"operational_compliance\":\"Not In Compliance\",\"secondary_compliance\":\"Not In Compliance\"}]}"},
		{"compliance by key", "/compliances/13", "{\"data\":{\"id\":13,\"report_id\":2,\"paragraph_id\":3,\"primary_compliance\":\"In Compliance\",\"operational_compliance\":\"Not In Compliance\",\"secondary_compliance\":\"Not In Compliance\"}}"},
		{"all reports", "/reports", "{\"data\":[{\"id\":1,\"report_name\":\"IMR-1\",\"report_title\":\"Monitor's First Report\",\"publish_date\":\"2015-11-23\",\"period_begin\":\"2015-02-01\",\"period_end\":\"2015-05-31\"}]}"},
		{"report by key", "/reports/1", "{\"data\":{\"id\":1,\"report_name\":\"IMR-1\",\"report_title\":\"Monitor's First Report\",\"publish_date\":\"2015-11-23\",\"period_begin\":\"2015-02-01\",\"period_end\":\"2015-05-31\"}}"},
	}
	for _, test := range tests {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", test.URL, nil)
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(rr, req)
		if test.expected != rr.Body.String() {
			t.Errorf("\n%v\n...expected = %v\n...obtained = %v", test.description, test.expected, rr.Body.String())
		}
	}
}
