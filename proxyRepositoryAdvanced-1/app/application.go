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
