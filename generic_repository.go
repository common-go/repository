package repository

import "context"

type GenericRepository interface {
	ViewRepository
	Insert(ctx context.Context, model interface{}) (int64, error)
	Update(ctx context.Context, model interface{}) (int64, error)
	Patch(ctx context.Context, model map[string]interface{}) (int64, error)
	Save(ctx context.Context, model interface{}) (int64, error)
	Delete(ctx context.Context, id interface{}) (int64, error)
}
