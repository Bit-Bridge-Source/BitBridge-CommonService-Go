package adapter

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoAdapter interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
}

type MongoAdapterImpl struct {
	Adapter MongoAdapter
}

// Find implements IUserMongoAdapter.
func (m *MongoAdapterImpl) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return m.Adapter.Find(ctx, filter, opts...)
}

// DeleteOne implements IUserMongoAdapter.
func (m *MongoAdapterImpl) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return m.Adapter.DeleteOne(ctx, filter, opts...)
}

// FindOne implements IUserMongoAdapter.
func (m *MongoAdapterImpl) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return m.Adapter.FindOne(ctx, filter)
}

// InsertOne implements IUserMongoAdapter.
func (m *MongoAdapterImpl) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return m.Adapter.InsertOne(ctx, document, opts...)
}

// UpdateOne implements IUserMongoAdapter.
func (m *MongoAdapterImpl) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return m.Adapter.UpdateOne(ctx, filter, update, opts...)
}

// NewMongoAdapter creates a new UserMongoAdapter
func NewMongoAdapter(adapter MongoAdapter) *MongoAdapterImpl {
	return &MongoAdapterImpl{
		Adapter: adapter,
	}
}

var _ MongoAdapter = (*MongoAdapterImpl)(nil)
