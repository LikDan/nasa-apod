package api

import (
	"context"
	"os"
)

type RepositoryCDN interface {
	GetByPath(ctx context.Context, path string) ([]byte, error)
}

type repositoryCDN struct {
	storagePath string
}

func NewRepositoryCDN(storagePath string) RepositoryCDN {
	return &repositoryCDN{
		storagePath: storagePath,
	}
}

func (r *repositoryCDN) GetByPath(_ context.Context, path string) ([]byte, error) {
	return os.ReadFile(r.storagePath + "/" + path)
}
