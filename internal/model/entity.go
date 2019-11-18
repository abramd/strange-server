package model

import (
	"time"
)

const EntityTable = "entities"
const SourceTable = "sources"

type State string

const (
	StateNew     = "new"
	StateOld     = "old"
	StateDeleted = "deleted"
)

type Entity struct {
	SourceName  string `json:"-"`
	Action      string `json:"action"`
	State       State  `json:"state"`
	IsProcessed bool   `json:"is_processed"`
}

type EntityPresenter struct {
	ID          int        `db:"id" json:"id"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
	SourceID    *int       `db:"source_id" json:"source_id"`
	Action      string     `db:"action" json:"action"`
	State       State      `db:"state" json:"state"`
	IsProcessed bool       `db:"is_processed" json:"is_processed"`
}
