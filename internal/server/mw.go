package server

import (
	"fmt"
	"github.com/abramd/strange-server/internal/model"
	"net/http"
	"sync"
)

const SourceHeader = "Source-Type"

var SourceHeaderError = fmt.Errorf("")

type SourceMap struct {
	mu   sync.RWMutex
	data map[string]int
}

func (m *SourceMap) Get(k string) (int, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	res, ok := m.data[k]
	return res, ok
}

func NewSourceMap(m *model.Manager) (*SourceMap, error) {
	ss, err := m.ListSources()
	if err != nil {
		return nil, err
	}

	data := make(map[string]int)
	for _, s := range ss {
		data[s.Name] = s.ID
	}
	return &SourceMap{
		mu:   sync.RWMutex{},
		data: data,
	}, nil
}

func SourceValidationMW(m *SourceMap, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, ok := m.Get(r.Header.Get(SourceHeader)); !ok {
			http.Error(w, SourceHeaderError.Error(), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}
}
