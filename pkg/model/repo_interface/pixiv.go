package repo_interface

import "context"

type PicRepo interface {
	FetchPixivPictureToMem(ctx context.Context, url string) ([]byte, error)
}
