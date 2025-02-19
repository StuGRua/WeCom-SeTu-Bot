package repo_interface

import (
	"context"
	"server/pkg/model/entity"
)

type SeTuRepo interface {
	GetArchiveInfoSlice(ctx context.Context, query *entity.Query) (result entity.QueryResult, err error)
}
