package dll


// PostRepositoryInit initial
func PostRepositoryInit() *PostRepo {
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

func test1() {
	a := PostRepositoryInit()
	a.GetByID(12)
}
