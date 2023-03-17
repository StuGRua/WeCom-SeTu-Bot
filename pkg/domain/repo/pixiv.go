package repo

import "context"

type PicRepo interface {
	FetchPixivPictureToMem(ctx context.Context, url string) ([]byte, error)
}
