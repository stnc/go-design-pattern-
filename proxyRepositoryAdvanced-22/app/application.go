package dll

import (
	"fmt"
	cach "proxy/app/cacheRepo"
	repo "proxy/app/standartRepo"
)

type PostAppInterface interface {
	GetByID(uint64) string
	GetBySelman(uint64) string
}

type Repositories struct {
	Post PostAppInterface
}

func postRepo() Repositories {
	repog := Repositories{
		Post: repo.PostRepositoryInit(),
	}
	return repog
}

func cacheRepo() Repositories {
	repog := Repositories{
		Post: cach.CACHERepositoriesInit(),
	}
	return repog
}

func Repos2() Repositories {

	var cacheType bool = true
	if cacheType == true {

		return cacheRepo()
	} else {
		return postRepo()
	}
}
func TheEnd() {
	fmt.Println(Repos2().Post.GetByID(12))
}
