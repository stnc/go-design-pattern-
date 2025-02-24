package proxy

type Application struct {
}

func (a *Application) response(cacheType bool) string {
	if cacheType == true {
		return a.cacheRepo()
	} else {
		return a.normalRepo()
	}

}
