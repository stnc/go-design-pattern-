package dll

import (
	cach "proxy/app/cacheRepo"
	repo "proxy/app/standartRepo"
	"fmt"
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

type PostAppInterface interface {
	GetByID(uint64) string
	GetBySelman(uint64) string
}

type Repositories struct {
	Post PostAppInterface
}

func RepositoriesInit() (*Repositories, error) {
	return &Repositories{
		Post: repo.PostRepositoryInit(),
	}, nil
}

func CACHERepositoriesInit() (*Repositories, error) {
	return &Repositories{
		Post: cach.CACHERepositoriesInit(),
	}, nil
}

func Repos() Repositories {
	repog := Repositories{
		Post: repo.PostRepositoryInit(),
	}
	return repog
}

func ReposC() Repositories {
	repog := Repositories{
		Post: cach.CACHERepositoriesInit(),
	}
	return repog
}

func Repos2() Repositories {

	var cacheType bool = true
	if cacheType == true {
		b := ReposC()
		b.Post.GetByID(12)
		return b
	} else {
		c := Repos()
		c.Post.GetByID(12)
		return c
	}
}
func TheEnd()  {
	fmt.Println(Repos2().Post.GetByID(12))
}

func Repos33() {

	var b *repo.PostRepo
	b = repo.PostRepositoryInit()
	b.GetByID(12)

	var c *cach.PostRepo
	c = cach.CACHERepositoriesInit()
	c.GetByID(12)
}
func Test3() {

	repo := Repositories{
		Post: repo.PostRepositoryInit(),
	}

	repo.Post.GetByID(12)
	Repos().Post.GetByID(12)

}
