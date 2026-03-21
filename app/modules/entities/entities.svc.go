package entities

import (
	"github.com/uptrace/bun"
)

type Service struct {
	db   bun.IDB
	root *bun.DB
}

func newService(db *bun.DB) *Service {
	return &Service{
		db:   db,
		root: db,
	}
}
