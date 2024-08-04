package repository

var repositories = func() map[string]*RepositoryFS { return make(map[string]*RepositoryFS) }()

func AddRepository(fs *RepositoryFS) {
	repositories[fs.name] = fs
}

func GetRepository(key string) *RepositoryFS {
	return repositories[key]
}
