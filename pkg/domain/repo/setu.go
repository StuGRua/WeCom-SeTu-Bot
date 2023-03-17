package repo

import (
	"context"
	"server/pkg/domain/entity"
)

type SetuRepo interface {
	GetArchiveInfoSlice(ctx context.Context, query *entity.Query) (result entity.QueryResult, err error)
}
