package repository

import (
	"context"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CRUD defines the contract for CRUD operations
type CRUD interface {
	Create(ctx context.Context, model mgm.Model) error
	GetByID(ctx context.Context, id string, result mgm.Model) error
	GetAll(ctx context.Context, filter bson.M, results interface{}, opts ...*options.FindOptions) error
	Update(ctx context.Context, id string, update bson.M, result mgm.Model) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context, filter bson.M) (int64, error)
	FindWithPagination(ctx context.Context, filter bson.M, page, limit int64, results interface{}) (int64, error)
	FindOne(ctx context.Context, filter bson.M, result mgm.Model) error
	UpdateMany(ctx context.Context, filter bson.M, update bson.M) (int64, error)
	DeleteMany(ctx context.Context, filter bson.M) (int64, error)
}
