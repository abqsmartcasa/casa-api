package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/abqsmartcasa/casa-api/models"

	"github.com/gorilla/context"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "max-age=3600")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (env *Env) paragraphs(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	ps, err := env.db.AllParagraphs(lang)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	responseJSON(w, ps, lang)
}

func (env *Env) paragraph(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	paragraph := models.Paragraph{}
	paragraph.ID = context.Get(r, "key").(int)
	p, err := env.db.GetParagraph(lang, paragraph)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	responseJSON(w, p, lang)
}

func (env *Env) paragraphsBySpecificTag(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	specificTag := models.SpecificTag{}
	specificTag.ID = context.Get(r, "key").(int)
	ps, err := env.db.GetParagraphsBySpecificTag(lang, specificTag)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	responseJSON(w, ps, lang)
}

func (env *Env) paragraphsByCategoryTag(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	categoryTag := models.CategoryTag{}
	categoryTag.ID = context.Get(r, "key").(int)
	ps, err := env.db.GetParagraphsByCategoryTag(lang, categoryTag)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	responseJSON(w, ps, lang)
}

func (env *Env) compliances(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	cs, err := env.db.AllCompliances(lang)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	responseJSON(w, cs, lang)
}

func (env *Env) compliancesByParagraph(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	paragraph := models.Paragraph{}
	paragraph.ID = context.Get(r, "key").(int)
	cs, err := env.db.GetCompliancesByParagraph(lang, paragraph)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	responseJSON(w, cs, lang)
}

func (env *Env) compliancesByReport(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	report := models.Report{}
	report.ID = context.Get(r, "key").(int)
	cs, err := env.db.GetCompliancesByReport(lang, report)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	responseJSON(w, cs, lang)
}

func (env *Env) reports(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	rpts, err := env.db.AllReports(lang)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	responseJSON(w, rpts, lang)
}

func (env *Env) report(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	report := models.Report{}
	report.ID = context.Get(r, "key").(int)
	rpt, err := env.db.GetReport(lang, report)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	responseJSON(w, rpt, lang)
}

func (env *Env) categoryTags(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	cts, err := env.db.AllCategoryTags(lang)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	responseJSON(w, cts, lang)
}

func (env *Env) categoryTag(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	categoryTag := models.CategoryTag{}
	categoryTag.ID = context.Get(r, "key").(int)
	ct, err := env.db.GetCategoryTag(lang, categoryTag)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	responseJSON(w, ct, lang)
}

func (env *Env) categoryTagsByParagraph(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	paragraph := models.Paragraph{}
	paragraph.ID = context.Get(r, "key").(int)
	cts, err := env.db.GetCategoryTagsByParagraph(lang, paragraph)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	responseJSON(w, cts, lang)
}

func (env *Env) specificTags(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	sts, err := env.db.AllSpecificTags(lang)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	responseJSON(w, sts, lang)
}

func (env *Env) specificTag(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	specificTag := models.SpecificTag{}
	specificTag.ID = context.Get(r, "key").(int)
	st, err := env.db.GetSpecificTag(lang, specificTag)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	responseJSON(w, st, lang)
}

func (env *Env) specificTagsByParagraph(w http.ResponseWriter, r *http.Request) {
	lang := context.Get(r, "lang")
	paragraph := models.Paragraph{}
	paragraph.ID = context.Get(r, "key").(int)
	sts, err := env.db.GetSpecificTagsByParagraph(lang, paragraph)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	responseJSON(w, sts, lang)
}
