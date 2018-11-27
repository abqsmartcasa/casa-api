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
	p := &models.Paragraph{}
	if paragraph.ID == 14 {
		p.ID = 14
		p.ParagraphNumber = 14
		p.ParagraphTitle = "test"
		p.ParagraphText = "test"
	}
	return p, nil
}

func (mdb *mockDB) GetParagraphsByCategoryTag(lang interface{}, categoryTag models.CategoryTag) ([]*models.Paragraph, error) {
	ps := make([]*models.Paragraph, 0)
	if categoryTag.ID == 1 {
		ps = append(ps, &models.Paragraph{
			ID:              42,
			ParagraphNumber: 42,
			ParagraphTitle:  "test",
			ParagraphText:   "test",
		})
	}
	return ps, nil
}

func (mdb *mockDB) GetParagraphsBySpecificTag(lang interface{}, specificTag models.SpecificTag) ([]*models.Paragraph, error) {
	ps := make([]*models.Paragraph, 0)
	if specificTag.ID == 1 {
		ps = append(ps, &models.Paragraph{
			ID:              42,
			ParagraphNumber: 42,
			ParagraphTitle:  "test",
			ParagraphText:   "test",
		})
	}
	return ps, nil
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
	c := &models.Compliance{}
	if compliance.ID == 13 {
		c.ID = 13
		c.ReportID = 2
		c.ParagraphID = 3
		c.PrimaryCompliance = "In Compliance"
		c.SecondaryCompliance = "Not In Compliance"
		c.OperationCompliance = "Not In Compliance"
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
	rpt := &models.Report{}
	if report.ID == 1 {
		rpt.ID = 1
		rpt.ReportName = "IMR-1"
		rpt.ReportTitle = "Monitor's First Report"
		rpt.PublishDate = "2015-11-23"
		rpt.PeriodBegin = "2015-02-01"
		rpt.PeriodEnd = "2015-05-31"
	}
	return rpt, nil
}

func (mdb *mockDB) AllCategoryTags(lang interface{}) ([]*models.CategoryTag, error) {
	cts := make([]*models.CategoryTag, 0)
	cts = append(cts, &models.CategoryTag{
		ID:    1,
		Value: "I. Use of Force",
	})
	return cts, nil
}

func (mdb *mockDB) GetCategoryTag(lang interface{}, categoryTag models.CategoryTag) (*models.CategoryTag, error) {
	ct := &models.CategoryTag{}
	if categoryTag.ID == 1 {
		ct.ID = 1
		ct.Value = "I. Use of Force"
	}
	return ct, nil
}

func (mdb *mockDB) GetCategoryTagsByParagraph(lang interface{}, paragraph models.Paragraph) ([]*models.CategoryTag, error) {
	cts := make([]*models.CategoryTag, 0)
	if paragraph.ID == 14 {
		cts = append(cts, &models.CategoryTag{
			ID:    1,
			Value: "I. Use of Force",
		})
	}
	return cts, nil
}

func (mdb *mockDB) AllSpecificTags(lang interface{}) ([]*models.SpecificTag, error) {
	sts := make([]*models.SpecificTag, 0)
	sts = append(sts, &models.SpecificTag{
		ID:         1,
		Value:      "Use of Force Principles",
		CategoryID: 1,
	})
	return sts, nil
}

func (mdb *mockDB) GetSpecificTag(lang interface{}, specificTag models.SpecificTag) (*models.SpecificTag, error) {
	st := &models.SpecificTag{}
	if specificTag.ID == 1 {
		st.ID = 1
		st.Value = "Use of Force Principles"
		st.CategoryID = 1
	}
	return st, nil
}

func (mdb *mockDB) GetSpecificTagsByParagraph(lang interface{}, paragraph models.Paragraph) ([]*models.SpecificTag, error) {
	sts := make([]*models.SpecificTag, 0)
	if paragraph.ID == 14 {
		sts = append(sts, &models.SpecificTag{
			ID:         1,
			Value:      "Use of Force Principles",
			CategoryID: 1,
		})
	}
	return sts, nil
}
func TestHandlers(t *testing.T) {
	router := mux.NewRouter()
	env := Env{db: &mockDB{}}
	router.HandleFunc("/paragraphs", env.paragraphs)
	router.HandleFunc("/paragraphs/{key}", env.paragraph)
	router.HandleFunc("/paragraphs/{key}/categorytags", env.categoryTagsByParagraph)
	router.HandleFunc("/paragraphs/{key}/specifictags", env.specificTagsByParagraph)
	router.HandleFunc("/compliances", env.compliances)
	router.HandleFunc("/compliances/{key}", env.compliance)
	router.HandleFunc("/reports", env.reports)
	router.HandleFunc("/reports/{key}", env.report)
	router.HandleFunc("/categorytags", env.categoryTags)
	router.HandleFunc("/categorytags/{key}", env.categoryTag)
	router.HandleFunc("/categorytags/{key}/paragraphs", env.paragraphsByCategoryTag)
	router.HandleFunc("/specifictags", env.specificTags)
	router.HandleFunc("/specifictags/{key}", env.specificTag)
	router.HandleFunc("/specifictags/{key}/paragraphs", env.paragraphsBySpecificTag)
	tests := []struct {
		description string
		Code        int
		URL         string
		expected    string
	}{
		{"all paragraphs", 200, "/paragraphs", "{\"data\":[{\"id\":42,\"paragraph_number\":42,\"paragraph_title\":\"test\",\"paragraph_text\":\"test\"}]}"},
		{"paragraph by key", 200, "/paragraphs/14", "{\"data\":{\"id\":14,\"paragraph_number\":14,\"paragraph_title\":\"test\",\"paragraph_text\":\"test\"}}"},
		{"invalid paragraph key", 404, "/paragraphs/13", ""},
		{"paragraphs by category tag", 200, "/categorytags/1/paragraphs", "{\"data\":[{\"id\":42,\"paragraph_number\":42,\"paragraph_title\":\"test\",\"paragraph_text\":\"test\"}]}"},
		{"paragraphs by invalid category tag", 404, "/categorytags/13/paragraphs", ""},
		{"paragraphs by specific tag", 200, "/specifictags/1/paragraphs", "{\"data\":[{\"id\":42,\"paragraph_number\":42,\"paragraph_title\":\"test\",\"paragraph_text\":\"test\"}]}"},
		{"paragraphs by invalid specific tag", 404, "/specifictags/13/paragraphs", ""},
		{"specific tags by invalid paragraph", 404, "/paragraphs/13/specifictags", ""},
		{"all compliances", 200, "/compliances", "{\"data\":[{\"id\":1,\"report_id\":2,\"paragraph_id\":3,\"primary_compliance\":\"In Compliance\",\"operational_compliance\":\"Not In Compliance\",\"secondary_compliance\":\"Not In Compliance\"}]}"},
		{"compliance by key", 200, "/compliances/13", "{\"data\":{\"id\":13,\"report_id\":2,\"paragraph_id\":3,\"primary_compliance\":\"In Compliance\",\"operational_compliance\":\"Not In Compliance\",\"secondary_compliance\":\"Not In Compliance\"}}"},
		{"invalid compliances key", 404, "/compliances/42", ""},
		{"all reports", 200, "/reports", "{\"data\":[{\"id\":1,\"report_name\":\"IMR-1\",\"report_title\":\"Monitor's First Report\",\"publish_date\":\"2015-11-23\",\"period_begin\":\"2015-02-01\",\"period_end\":\"2015-05-31\"}]}"},
		{"report by key", 200, "/reports/1", "{\"data\":{\"id\":1,\"report_name\":\"IMR-1\",\"report_title\":\"Monitor's First Report\",\"publish_date\":\"2015-11-23\",\"period_begin\":\"2015-02-01\",\"period_end\":\"2015-05-31\"}}"},
		{"invalid report key", 404, "/reports/42", ""},
		{"all category tags", 200, "/categorytags", "{\"data\":[{\"id\":1,\"value\":\"I. Use of Force\"}]}"},
		{"category tag by key", 200, "/categorytags/1", "{\"data\":{\"id\":1,\"value\":\"I. Use of Force\"}}"},
		{"invalid category tag key", 404, "/categorytags/42", ""},
		{"category tags by paragraph", 200, "/paragraphs/14/categorytags", "{\"data\":[{\"id\":1,\"value\":\"I. Use of Force\"}]}"},
		{"category tags by invalid paragraph", 404, "/paragraphs/13/categorytags", ""},
		{"all specific tags", 200, "/specifictags", "{\"data\":[{\"id\":1,\"value\":\"Use of Force Principles\",\"category_id\":1}]}"},
		{"specific tag by key", 200, "/specifictags/1", "{\"data\":{\"id\":1,\"value\":\"Use of Force Principles\",\"category_id\":1}}"},
		{"invalid specific tag key", 404, "/specifictags/42", ""},
		{"specific tags by paragraph", 200, "/paragraphs/14/specifictags", "{\"data\":[{\"id\":1,\"value\":\"Use of Force Principles\",\"category_id\":1}]}"},
		{"specific tags by invalid paragraph", 404, "/paragraphs/13/specifictags", ""},
	}
	for _, test := range tests {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", test.URL, nil)
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(rr, req)
		if test.Code != rr.Code {
			t.Errorf("\n%v\nhandler returned wrong status code: got %v want %v",
				test.description, rr.Code, test.Code)
		}
		if test.expected != rr.Body.String() && rr.Code == http.StatusOK {
			t.Errorf("\n%v\n...expected = %v\n...obtained = %v",
				test.description, test.expected, rr.Body.String())
		}
	}
}
