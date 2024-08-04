package files

import (
	"errors"
	"fmt"
	"io"
	"io/fs"

	"github.com/NublyBR/go-vpk"
)

type VpkFS struct {
	paths []string
	vpks  []vpk.VPK
}

func NewVpkFS(paths ...string) *VpkFS {
	vpk := VpkFS{
		paths: paths,
		vpks:  make([]vpk.VPK, 0, 1),
	}

	vpk.init()

	return &vpk
}

func (fs *VpkFS) init() error {
	for _, path := range fs.paths {
		vpk, err := vpk.OpenDir(path)

		if err != nil {
			return err
		}

		fs.vpks = append(fs.vpks, vpk)
	}

	return nil
}

func (fs *VpkFS) Open(name string) (fs.File, error) {
	return nil, nil
}

func (fs *VpkFS) ReadFile(path string) ([]byte, error) {
	var entry vpk.Entry
	var ok bool
	for _, vpk := range fs.vpks {
		if entry, ok = vpk.Find(path); ok {
			break
		}
	}

	if !ok {
		return nil, errors.New("file not found")
	}

	fileReader, err := entry.Open()
	if err != nil {
		return nil, fmt.Errorf("unable to open file: <%w>", err)
	}

	buf, err := io.ReadAll(fileReader)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: <%w>", err)
	}

	return buf, nil
}
