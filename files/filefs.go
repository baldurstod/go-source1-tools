package files

import (
	"io/fs"
	"os"
	"path"
)

type FileFS struct {
	path string
}

func NewFileFS(path string) *FileFS {
	return &FileFS{
		path: path,
	}
}

func (fs *FileFS) Open(name string) (fs.File, error) {
	return nil, nil
}

func (fs *FileFS) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(path.Join(fs.path, name))
}
