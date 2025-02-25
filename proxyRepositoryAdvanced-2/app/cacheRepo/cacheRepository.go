package dll

type Cache struct {
}

func (c Cache) GetByUsername() string {
	return "cache"
}

func TesGetByUsername() string {
	return "cache"
}

///-----

// PostRepositoryInit initial
func CACHERepositoriesInit() *PostRepo {
	return &PostRepo{}
}

type PostRepo struct {
}

func (r *PostRepo) GetByID(id uint64) string {
	return "data, GetByID"
}

func (r *PostRepo) GetBySelman(id uint64) string {
	return "data GetBySelman"
}

