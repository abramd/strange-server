package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Manager struct {
	db *sqlx.DB
}

func NewManager(db *sqlx.DB) *Manager {
	return &Manager{
		db: db,
	}
}

func (m *Manager) ExecTx(query string, args ...interface{}) error {
	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, args...)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (m *Manager) Insert(e *Entity) error {
	_, err := m.db.Exec(fmt.Sprintf(
		`INSERT INTO %s (action, state, source_id) 
				VALUES ($1,$2,(SELECT id FROM %s WHERE name = $3 LIMIT 1))`,
		EntityTable, SourceTable),
		e.Action,
		e.State,
		e.SourceName)
	return err
}

func (m *Manager) ListEntities() ([]*EntityPresenter, error) {
	rr, err := m.db.Queryx("SELECT * FROM entities")
	if err != nil {
		return nil, err
	}
	res := make([]*EntityPresenter, 0)
	for rr.Next() {
		s := new(EntityPresenter)
		err := rr.StructScan(s)
		if err != nil {
			return nil, err
		}
		res = append(res, s)
	}

	return res, nil
}

func (m *Manager) ListSources() ([]*SourcePresenter, error) {
	rr, err := m.db.Queryx("SELECT * FROM sources")
	if err != nil {
		return nil, err
	}
	res := make([]*SourcePresenter, 0)
	for rr.Next() {
		s := new(SourcePresenter)
		err := rr.StructScan(s)
		if err != nil {
			return nil, err
		}
		res = append(res, s)
	}

	return res, nil
}
