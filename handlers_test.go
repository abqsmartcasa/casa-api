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

func TestParagraphs(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/paragraphs", nil)
	if err != nil {
		t.Fatal(err)
	}
	router := mux.NewRouter()
	env := Env{db: &mockDB{}}
	router.HandleFunc("/paragraphs", env.paragraphs)
	router.ServeHTTP(rr, req)
	expected := "{\"data\":[{\"id\":42,\"paragraph_number\":42,\"paragraph_title\":\"test\",\"paragraph_text\":\"test\"}]}"
	if expected != rr.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rr.Body.String())
	}
}

func TestParagraph(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/paragraphs/13", nil)
	if err != nil {
		t.Fatal(err)
	}
	router := mux.NewRouter()
	env := Env{db: &mockDB{}}
	router.HandleFunc("/paragraphs/{key}", env.paragraph)
	router.ServeHTTP(rr, req)
	expected := "{\"data\":{\"id\":13,\"paragraph_number\":13,\"paragraph_title\":\"test\",\"paragraph_text\":\"test\"}}"
	if expected != rr.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rr.Body.String())
	}
}

func TestCompliances(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/compliances", nil)
	if err != nil {
		t.Fatal(err)
	}
	router := mux.NewRouter()
	env := Env{db: &mockDB{}}
	router.HandleFunc("/compliances", env.compliances)
	router.ServeHTTP(rr, req)
	expected := "{\"data\":[{\"id\":1,\"report_id\":2,\"paragraph_id\":3,\"primary_compliance\":\"In Compliance\",\"operational_compliance\":\"Not In Compliance\",\"secondary_compliance\":\"Not In Compliance\"}]}"
	if expected != rr.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rr.Body.String())
	}
}

func TestCompliance(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/compliances/13", nil)
	if err != nil {
		t.Fatal(err)
	}
	router := mux.NewRouter()
	env := Env{db: &mockDB{}}
	router.HandleFunc("/compliances/{key}", env.compliance)
	router.ServeHTTP(rr, req)
	expected := "{\"data\":{\"id\":13,\"report_id\":2,\"paragraph_id\":3,\"primary_compliance\":\"In Compliance\",\"operational_compliance\":\"Not In Compliance\",\"secondary_compliance\":\"Not In Compliance\"}}"
	if expected != rr.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rr.Body.String())
	}
}

func TestReports(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/reports", nil)
	if err != nil {
		t.Fatal(err)
	}
	router := mux.NewRouter()
	env := Env{db: &mockDB{}}
	router.HandleFunc("/reports", env.reports)
	router.ServeHTTP(rr, req)
	expected := "{\"data\":[{\"id\":1,\"report_name\":\"IMR-1\",\"report_title\":\"Monitor's First Report\",\"publish_date\":\"2015-11-23\",\"period_begin\":\"2015-02-01\",\"period_end\":\"2015-05-31\"}]}"
	if expected != rr.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rr.Body.String())
	}
}

func TestReport(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/reports/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	router := mux.NewRouter()
	env := Env{db: &mockDB{}}
	router.HandleFunc("/reports/{key}", env.report)
	router.ServeHTTP(rr, req)
	expected := "{\"data\":{\"id\":1,\"report_name\":\"IMR-1\",\"report_title\":\"Monitor's First Report\",\"publish_date\":\"2015-11-23\",\"period_begin\":\"2015-02-01\",\"period_end\":\"2015-05-31\"}}"
	if expected != rr.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rr.Body.String())
	}
}
