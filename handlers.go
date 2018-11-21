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

type responseData struct {
	Data json.RawMessage `json:"data"`
}

func (env *Env) paragraphs(w http.ResponseWriter, r *http.Request) {

	lang := context.Get(r, "lang")
	paragraphs, err := env.db.AllParagraphs(lang)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	psJSON, err := json.Marshal(paragraphs)
	res := responseData{psJSON}
	data, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "Content-Type: application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
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
	pJSON, err := json.Marshal(p)
	res := responseData{pJSON}
	data, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "Content-Type: application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
