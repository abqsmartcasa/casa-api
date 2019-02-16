package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apdforward/apdf_api/models"
	"github.com/gorilla/mux"
)

type Pages struct {
	Message []int
}

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

func (mdb *mockDB) GetParagraph(lang interface{}, paragraph models.Paragraph) (*models.Paragraph, error) {
	p := &models.Paragraph{}
	if paragraph.ID == 14 {
		p.ID = 14
		p.ParagraphNumber = 14
		p.ParagraphTitle = "test"
		p.ParagraphText = "test"
		return p, nil
	}
	err := errors.New("No such paragraph")
	return nil, err

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
		return ps, nil
	}
	err := errors.New("No such paragraph")
	return nil, err

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
		return ps, nil
	}
	err := errors.New("No such paragraph")
	return nil, err
}

func (mdb *mockDB) AllCompliances(lang interface{}) ([]*models.Compliance, error) {
	cs := make([]*models.Compliance, 0)
	pages := []int{14, 15}
	byte1, err := json.Marshal(pages)
	if err != nil {
		fmt.Println(err)
	}
	raw := json.RawMessage(byte1)
	cs = append(cs, &models.Compliance{
		ReportID:            2,
		ParagraphID:         3,
		PrimaryCompliance:   "In Compliance",
		SecondaryCompliance: "Not In Compliance",
		OperationCompliance: "Not In Compliance",
		Pages:               raw,
	})
	return cs, nil
}

func (mdb *mockDB) GetCompliancesByParagraph(lang interface{}, paragraph models.Paragraph) ([]*models.Compliance, error) {
	cs := make([]*models.Compliance, 0)
	pages := []int{14, 15}
	byte1, err := json.Marshal(pages)
	if err != nil {
		fmt.Println(err)
	}
	raw := json.RawMessage(byte1)
	if paragraph.ID == 14 {
		cs = append(cs, &models.Compliance{
			ReportID:            2,
			ParagraphID:         14,
			PrimaryCompliance:   "In Compliance",
			SecondaryCompliance: "Not In Compliance",
			OperationCompliance: "Not In Compliance",
			Pages:               raw,
		})
		return cs, nil
	}
	err = errors.New("No such compliances")
	return nil, err

}

func (mdb *mockDB) GetCompliancesByReport(lang interface{}, report models.Report) ([]*models.Compliance, error) {
	cs := make([]*models.Compliance, 0)
	pages := []int{14, 15}
	byte1, err := json.Marshal(pages)
	if err != nil {
		fmt.Println(err)
	}
	raw := json.RawMessage(byte1)
	if report.ID == 2 {
		cs = append(cs, &models.Compliance{
			ReportID:            2,
			ParagraphID:         14,
			PrimaryCompliance:   "In Compliance",
			SecondaryCompliance: "Not In Compliance",
			OperationCompliance: "Not In Compliance",
			Pages:               raw,
		})
		return cs, nil
	}
	err = errors.New("No such compliances")
	return nil, err
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
		return rpt, nil
	}
	err := errors.New("No such report")
	return nil, err
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
		return ct, nil
	}
	err := errors.New("No such category tag")
	return nil, err
}

func (mdb *mockDB) GetCategoryTagsByParagraph(lang interface{}, paragraph models.Paragraph) ([]*models.CategoryTag, error) {
	cts := make([]*models.CategoryTag, 0)
	if paragraph.ID == 14 {
		cts = append(cts, &models.CategoryTag{
			ID:    1,
			Value: "I. Use of Force",
		})
		return cts, nil
	}
	err := errors.New("invalid paragraph")
	return nil, err
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
		return st, nil
	}
	err := errors.New("No such specific tag")
	return nil, err
}

func (mdb *mockDB) GetSpecificTagsByParagraph(lang interface{}, paragraph models.Paragraph) ([]*models.SpecificTag, error) {
	sts := make([]*models.SpecificTag, 0)
	if paragraph.ID == 14 {
		sts = append(sts, &models.SpecificTag{
			ID:         1,
			Value:      "Use of Force Principles",
			CategoryID: 1,
		})
		return sts, nil
	}
	err := errors.New("invalid paragraph")
	return nil, err
}
func TestHandlers(t *testing.T) {
	router := mux.NewRouter()
	router.Use(handleKey)
	env := Env{db: &mockDB{}}
	router.HandleFunc("/paragraphs", env.paragraphs)
	router.HandleFunc("/paragraphs/{key}", env.paragraph)
	router.HandleFunc("/paragraphs/{key}/category-tags", env.categoryTagsByParagraph)
	router.HandleFunc("/paragraphs/{key}/specific-tags", env.specificTagsByParagraph)
	router.HandleFunc("/paragraphs/{key}/compliances", env.compliancesByParagraph)
	router.HandleFunc("/compliances", env.compliances)
	router.HandleFunc("/reports", env.reports)
	router.HandleFunc("/reports/{key}", env.report)
	router.HandleFunc("/reports/{key}/compliances", env.compliancesByReport)
	router.HandleFunc("/category-tags", env.categoryTags)
	router.HandleFunc("/category-tags/{key}", env.categoryTag)
	router.HandleFunc("/category-tags/{key}/paragraphs", env.paragraphsByCategoryTag)
	router.HandleFunc("/specific-tags", env.specificTags)
	router.HandleFunc("/specific-tags/{key}", env.specificTag)
	router.HandleFunc("/specific-tags/{key}/paragraphs", env.paragraphsBySpecificTag)
	tests := []struct {
		description string
		Code        int
		URL         string
		expected    string
	}{
		{"all paragraphs", 200, "/paragraphs", "{\"data\":[{\"id\":42,\"paragraphNumber\":42,\"paragraphTitle\":\"test\",\"paragraphText\":\"test\"}]}"},
		{"paragraph by key", 200, "/paragraphs/14", "{\"data\":{\"id\":14,\"paragraphNumber\":14,\"paragraphTitle\":\"test\",\"paragraphText\":\"test\"}}"},
		{"invalid paragraph key", 404, "/paragraphs/13", ""},
		{"paragraphs by category tag", 200, "/category-tags/1/paragraphs", "{\"data\":[{\"id\":42,\"paragraphNumber\":42,\"paragraphTitle\":\"test\",\"paragraphText\":\"test\"}]}"},
		{"paragraphs by invalid category tag", 404, "/category-tags/13/paragraphs", ""},
		{"paragraphs by specific tag", 200, "/specific-tags/1/paragraphs", "{\"data\":[{\"id\":42,\"paragraphNumber\":42,\"paragraphTitle\":\"test\",\"paragraphText\":\"test\"}]}"},
		{"paragraphs by invalid specific tag", 404, "/specific-tags/13/paragraphs", ""},
		{"specific tags by invalid paragraph", 404, "/paragraphs/13/specifictags", ""},
		{"all compliances", 200, "/compliances", "{\"data\":[{\"reportId\":2,\"paragraphId\":3,\"primaryCompliance\":\"In Compliance\",\"operationalCompliance\":\"Not In Compliance\",\"secondaryCompliance\":\"Not In Compliance\",\"pages\":[14,15]}]}"},
		{"invalid compliances key", 404, "/compliances/42", ""},
		{"compliances by paragraph", 200, "/paragraphs/14/compliances", "{\"data\":[{\"reportId\":2,\"paragraphId\":14,\"primaryCompliance\":\"In Compliance\",\"operationalCompliance\":\"Not In Compliance\",\"secondaryCompliance\":\"Not In Compliance\",\"pages\":[14,15]}]}"},
		{"compliances by invalid paragraph", 404, "/paragraphs/13/compliances", ""},
		{"compliances by report", 200, "/reports/2/compliances", "{\"data\":[{\"reportId\":2,\"paragraphId\":14,\"primaryCompliance\":\"In Compliance\",\"operationalCompliance\":\"Not In Compliance\",\"secondaryCompliance\":\"Not In Compliance\",\"pages\":[14,15]}]}"},
		{"compliances by invalid report", 404, "/reports/1/compliances", ""},
		{"all reports", 200, "/reports", "{\"data\":[{\"id\":1,\"reportName\":\"IMR-1\",\"reportTitle\":\"Monitor's First Report\",\"publishDate\":\"2015-11-23\",\"periodBegin\":\"2015-02-01\",\"periodEnd\":\"2015-05-31\"}]}"},
		{"report by key", 200, "/reports/1", "{\"data\":{\"id\":1,\"reportName\":\"IMR-1\",\"reportTitle\":\"Monitor's First Report\",\"publishDate\":\"2015-11-23\",\"periodBegin\":\"2015-02-01\",\"periodEnd\":\"2015-05-31\"}}"},
		{"invalid report key", 404, "/reports/42", ""},
		{"all category tags", 200, "/category-tags", "{\"data\":[{\"id\":1,\"value\":\"I. Use of Force\"}]}"},
		{"category tag by key", 200, "/category-tags/1", "{\"data\":{\"id\":1,\"value\":\"I. Use of Force\"}}"},
		{"invalid category tag key", 404, "/category-tags/42", ""},
		{"category tags by paragraph", 200, "/paragraphs/14/category-tags", "{\"data\":[{\"id\":1,\"value\":\"I. Use of Force\"}]}"},
		{"category tags by invalid paragraph", 404, "/paragraphs/13/category-tags", ""},
		{"all specific tags", 200, "/specific-tags", "{\"data\":[{\"id\":1,\"value\":\"Use of Force Principles\",\"categoryId\":1}]}"},
		{"specific tag by key", 200, "/specific-tags/1", "{\"data\":{\"id\":1,\"value\":\"Use of Force Principles\",\"categoryId\":1}}"},
		{"invalid specific tag key", 404, "/specific-tags/42", ""},
		{"specific tags by paragraph", 200, "/paragraphs/14/specific-tags", "{\"data\":[{\"id\":1,\"value\":\"Use of Force Principles\",\"categoryId\":1}]}"},
		{"specific tags by invalid paragraph", 404, "/paragraphs/13/specific-tags", ""},
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
