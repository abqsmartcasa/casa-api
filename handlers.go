package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/apdforward/apdf_api/models"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type Env struct {
	db models.Datastore
}

func responseJSON(w http.ResponseWriter, payload interface{}) {

	type responseData struct {
		Data json.RawMessage `json:"data"`
	}

	JSON, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	response := responseData{JSON}
	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "Content-Type: application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (env *Env) paragraphs(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	paragraphs, err := env.db.AllParagraphs(lang)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	responseJSON(w, paragraphs)
}

func (env *Env) paragraph(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	v := r.URL.Query()
	include := v.Get("include")
	vars := mux.Vars(r)
	paragraph := models.Paragraph{}
	paragraphID, err := strconv.Atoi(vars["key"])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	paragraph.ID = paragraphID
	p, err := env.db.GetParagraph(lang, paragraph, include)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	responseJSON(w, p)
}

func (env *Env) compliances(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	compliances, err := env.db.AllCompliances(lang)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	responseJSON(w, compliances)
}

func (env *Env) compliance(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	vars := mux.Vars(r)
	compliance := models.Compliance{}
	complianceID, err := strconv.Atoi(vars["key"])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	compliance.ID = complianceID
	c, err := env.db.GetCompliance(lang, compliance)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	responseJSON(w, c)
}

func (env *Env) reports(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	reports, err := env.db.AllReports(lang)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	responseJSON(w, reports)
}

func (env *Env) report(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	vars := mux.Vars(r)
	report := models.Report{}
	reportID, err := strconv.Atoi(vars["key"])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	report.ID = reportID
	rpt, err := env.db.GetReport(lang, report)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	responseJSON(w, rpt)
}

func (env *Env) categoryTags(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	categoryTags, err := env.db.AllCategoryTags(lang)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	responseJSON(w, categoryTags)
}

func (env *Env) categoryTag(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	vars := mux.Vars(r)
	categoryTag := models.CategoryTag{}
	categoryTagID, err := strconv.Atoi(vars["key"])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	categoryTag.ID = categoryTagID
	ct, err := env.db.GetCategoryTag(lang, categoryTag)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	responseJSON(w, ct)
}

func (env *Env) specificTags(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	specificTags, err := env.db.AllSpecificTags(lang)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	responseJSON(w, specificTags)
}

func (env *Env) specificTag(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	vars := mux.Vars(r)
	specificTag := models.SpecificTag{}
	specificTagID, err := strconv.Atoi(vars["key"])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	specificTag.ID = specificTagID
	st, err := env.db.GetSpecificTag(lang, specificTag)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	responseJSON(w, st)
}
