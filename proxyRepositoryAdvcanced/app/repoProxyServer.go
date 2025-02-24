package proxy

type Repo struct {
	application *Application
	manuelOpen  bool
	//	rateLimiter map[string]int
}

func RepoProxyServer() *Repo {
	return &Repo{
		application: &Application{},
		manuelOpen:  false,
		//	rateLimiter: make(map[string]int),
	}
}

func (n *Repo) RepoRequest() string {
	cacheType := n.check()

	return n.application.response(cacheType)
}

func (n *Repo) check() bool {
	cacheOK := false
	////override cache
	if n.manuelOpen == true {
		return true
	} else {
		if cacheOK == true {
	
			return true
		} else {

			return false
		}
	}

}
