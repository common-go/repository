package repository

import "context"

type ViewRepository interface {
	Ids() []string
	All(ctx context.Context) (interface{}, error)
	Load(ctx context.Context, id interface{}) (interface{}, error)
	LoadAndDecode(ctx context.Context, id interface{}, result interface{}) (bool, error)
	Exist(ctx context.Context, id interface{}) (bool, error)
}
