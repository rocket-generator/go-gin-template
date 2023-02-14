package seeds

import (
	"context"
	"github.com/uptrace/bun"
)

// Seed ... Struct for database seeding
type Seed struct {
	database *bun.DB
}

// Execute ...
func (seed *Seed) Execute(context context.Context) error {
	return nil
}

// NewSeed ... Create New Seed Instance
func NewSeed(db *bun.DB) *Seed {
	return &Seed{database: db}
}
