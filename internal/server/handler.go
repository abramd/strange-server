package server

import (
	"encoding/json"
	"github.com/abramd/strange-server/internal/model"
	"net/http"
)

var (
	SuccessBody = []byte("success")
)

func PostHandler(m model.EntityManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		entity := new(model.Entity)

		err := json.NewDecoder(r.Body).Decode(entity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		entity.SourceName = r.Header.Get(SourceHeader)

		err = m.Insert(entity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, _ = w.Write(SuccessBody)
	}
}

func ListHandler(m *model.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ee, err := m.ListEntities()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		b, err := json.Marshal(&ee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		_, _ = w.Write(b)
	}
}

func SourceListHandler(m *model.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ee, err := m.ListSources()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		b, err := json.Marshal(&ee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		_, _ = w.Write(b)
	}
}
