package repository

import (
	"errors"
	"io/fs"
)

type RepositoryFS struct {
	name        string
	filesystems []fs.ReadFileFS
}

func NewRepositoryFS(name string, filesystems ...fs.ReadFileFS) *RepositoryFS {
	repo := RepositoryFS{
		name:        name,
		filesystems: filesystems,
	}

	return &repo
}

func (fs *RepositoryFS) Open(name string) (fs.File, error) {
	return nil, nil
}

func (fs *RepositoryFS) ReadFile(path string) ([]byte, error) {
	var err error
	var b []byte
	for _, system := range fs.filesystems {

		if b, err = system.ReadFile(path); err == nil {
			return b, nil
		}
	}

	return nil, errors.New("file not found")
	/*
	   fileReader, err := entry.Open()

	   	if err != nil {
	   		return nil, fmt.Errorf("unable to open file: <%w>", err)
	   	}

	   buf, err := io.ReadAll(fileReader)

	   	if err != nil {
	   		return nil, fmt.Errorf("unable to read file: <%w>", err)
	   	}

	   return buf, nil
	*/
}

/*

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
*/
