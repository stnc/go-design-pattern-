package dll

// PostRepositoryInit initial
func CACHERepositoriesInit() *PostRepo {
	return &PostRepo{}
}

type PostRepo struct {
}

func (r *PostRepo) GetByID(id uint64) string {
	return "cache data, GetByID"
}

func (r *PostRepo) GetBySelman(id uint64) string {
	return "cache data GetBySelman"
}
