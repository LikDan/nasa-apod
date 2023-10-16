package updater

import (
	"io"
	"net/http"
	"os"
)

type RepositoryCDN interface {
	SaveFromURL(url string, path string) error
}

type repositoryCDN struct {
	storagePath string
}

func NewRepositoryCDN(storagePath string) RepositoryCDN {
	return &repositoryCDN{storagePath: storagePath}
}

func (r *repositoryCDN) SaveFromURL(url string, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	path = r.storagePath + "/" + path
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
