package proxy

type Cache struct {
}

type Standart struct {
}

type Application struct {
	Cache Cache
	Standart Standart
}

func (a *Application) response(cacheType bool) string {
	if cacheType == true {
		return a.Cache.GetByUsername()
	} else {
		return a.Standart.GetByUsername()
	}
}
