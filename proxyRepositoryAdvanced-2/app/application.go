package dll

import (
	cach "proxy/app/cacheRepo"
	repo "proxy/app/standartRepo"
)

type Application struct {
	Cache    cach.Cache
	Standart repo.Standart
}

func (a *Application) response(cacheType bool) string {
	if cacheType == true {
		return a.Cache.GetByUsername()
	} else {
		return a.Standart.GetByUsername()
	}
}

///--------
///--------
///--------
///--------

type PostAppInterface interface {
	GetByID(uint64) string
	GetBySelman(uint64) string
}

type Repositories struct {
	Post PostAppInterface
}

func RepositoriesInit() (*Repositories, error) {
	return &Repositories{

		Post: PostRepositoryInit(),
	}, nil
}

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
