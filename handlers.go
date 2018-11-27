package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/apdforward/apdf_api/models"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type Env struct {
	db models.Datastore
}

func responseJSON(w http.ResponseWriter, payload interface{}, lang interface{}) {

	type responseData struct {
		Data json.RawMessage `json:"data"`
	}

	JSON, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	response := responseData{JSON}
	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Language", fmt.Sprintf("%s", lang))
	//w.Header().Set("Cache-Control", "max-age=3600")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (env *Env) paragraphs(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	paragraphs, err := env.db.AllParagraphs(lang)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	responseJSON(w, paragraphs, lang)
}

func (env *Env) paragraph(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	v := r.URL.Query()
	include := v.Get("include")
	vars := mux.Vars(r)
	paragraph := models.Paragraph{}
	paragraphID, err := strconv.Atoi(vars["key"])
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	paragraph.ID = paragraphID
	p, err := env.db.GetParagraph(lang, paragraph, include)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if p.ID == 0 {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	responseJSON(w, p, lang)
}

func (env *Env) paragraphsBySpecificTag(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	vars := mux.Vars(r)
	specificTag := models.SpecificTag{}
	specificTagID, err := strconv.Atoi(vars["key"])
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	specificTag.ID = specificTagID
	ps, err := env.db.GetParagraphsBySpecificTag(lang, specificTag)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if ps == nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	responseJSON(w, ps, lang)
}

func (env *Env) paragraphsByCategoryTag(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	vars := mux.Vars(r)
	categoryTag := models.CategoryTag{}
	categoryTagID, err := strconv.Atoi(vars["key"])
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	categoryTag.ID = categoryTagID
	ps, err := env.db.GetParagraphsByCategoryTag(lang, categoryTag)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if ps == nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	responseJSON(w, ps, lang)
}

func (env *Env) compliances(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	compliances, err := env.db.AllCompliances(lang)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	responseJSON(w, compliances, lang)
}

func (env *Env) compliance(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	vars := mux.Vars(r)
	compliance := models.Compliance{}
	complianceID, err := strconv.Atoi(vars["key"])
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	compliance.ID = complianceID
	c, err := env.db.GetCompliance(lang, compliance)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if c.ID == 0 {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	responseJSON(w, c, lang)
}

func (env *Env) reports(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	reports, err := env.db.AllReports(lang)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	responseJSON(w, reports, lang)
}

func (env *Env) report(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	vars := mux.Vars(r)
	report := models.Report{}
	reportID, err := strconv.Atoi(vars["key"])
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	report.ID = reportID
	rpt, err := env.db.GetReport(lang, report)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if rpt.ID == 0 {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	responseJSON(w, rpt, lang)
}

func (env *Env) categoryTags(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	categoryTags, err := env.db.AllCategoryTags(lang)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	responseJSON(w, categoryTags, lang)
}

func (env *Env) categoryTag(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	vars := mux.Vars(r)
	categoryTag := models.CategoryTag{}
	categoryTagID, err := strconv.Atoi(vars["key"])
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	categoryTag.ID = categoryTagID
	ct, err := env.db.GetCategoryTag(lang, categoryTag)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if ct.ID == 0 {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	responseJSON(w, ct, lang)
}

func (env *Env) categoryTagsByParagraph(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	vars := mux.Vars(r)
	paragraph := models.Paragraph{}
	paragraphID, err := strconv.Atoi(vars["key"])
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	paragraph.ID = paragraphID
	cts, err := env.db.GetCategoryTagsByParagraph(lang, paragraph)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if cts == nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	responseJSON(w, cts, lang)
}

func (env *Env) specificTags(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	specificTags, err := env.db.AllSpecificTags(lang)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	responseJSON(w, specificTags, lang)
}

func (env *Env) specificTag(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	vars := mux.Vars(r)
	specificTag := models.SpecificTag{}
	specificTagID, err := strconv.Atoi(vars["key"])
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	specificTag.ID = specificTagID
	st, err := env.db.GetSpecificTag(lang, specificTag)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if st.ID == 0 {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	responseJSON(w, st, lang)
}

func (env *Env) specificTagsByParagraph(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	vars := mux.Vars(r)
	paragraph := models.Paragraph{}
	paragraphID, err := strconv.Atoi(vars["key"])
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	paragraph.ID = paragraphID
	sts, err := env.db.GetSpecificTagsByParagraph(lang, paragraph)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if sts == nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	responseJSON(w, sts, lang)
}
