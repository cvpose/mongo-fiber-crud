package repository

import (
	"context"

	"github.com/kamva/mgm/v3"
)

// Create inserts a new document into the collection
func (r *Repository) Create(ctx context.Context, model mgm.Model) error {
	return mgm.Coll(model).Create(model)
}
