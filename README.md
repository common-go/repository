# Repository
Some standard interfaces:
- StringRepository
- VerifiedCodeRepository
- ViewRepository: to get all data, to get data by identity
- GenericRepository: to insert, update, save, delete data

## Installation

Please make sure to initialize a Go module before installing common-go/repository:

```shell
go get -u github.com/common-go/repository
```

Import:

```go
import "github.com/common-go/repository"
```

You can optimize the import by version:
- v0.0.3: StringRepository
- v0.0.4: VerifiedCodeRepository
- v0.0.5: StringRepository, VerifiedCodeRepository, 
- v0.0.6: ViewRepository
- v0.0.8: ViewRepository, GenericRepository
- v0.0.9: ViewRepository, GenericRepository, StringRepository, VerifiedCodeRepository

## Details:
#### string_repository.go
```go
type StringRepository interface {
	Load(ctx context.Context, key string, max int64) ([]string, error)
	Save(ctx context.Context, values []string) (int64, error)
	Delete(ctx context.Context, values []string) (int64, error)
}
```

#### verified_code_repository.go
```go
type VerifiedCodeRepository interface {
	Save(ctx context.Context, id string, code string, expireAt time.Time) (int64, error)
	Load(ctx context.Context, id string) (string, time.Time, error)
	Delete(ctx context.Context, id string) (int64, error)
}
```

#### view_repository.go
```go
type ViewRepository interface {
	Ids() []string
	All(ctx context.Context) (interface{}, error)
	Load(ctx context.Context, id interface{}) (interface{}, error)
	LoadAndDecode(ctx context.Context, id interface{}, result interface{}) (bool, error)
	Exist(ctx context.Context, id interface{}) (bool, error)
}
```

#### generic_repository.go
```go
type GenericRepository interface {
	ViewRepository
	Insert(ctx context.Context, model interface{}) (int64, error)
	Update(ctx context.Context, model interface{}) (int64, error)
	Patch(ctx context.Context, model map[string]interface{}) (int64, error)
	Save(ctx context.Context, model interface{}) (int64, error)
	Delete(ctx context.Context, id interface{}) (int64, error)
}
```
