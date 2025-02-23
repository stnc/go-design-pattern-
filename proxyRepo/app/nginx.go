package proxy



type Nginx struct {
	application *Application
	manuelOpen  bool
	//	rateLimiter map[string]int
}

func NginxServer() *Nginx {
	return &Nginx{
		application: &Application{},
		manuelOpen:  false,
		//	rateLimiter: make(map[string]int),
	}
}

func (n *Nginx) RepoRequest() string {
	cacheType := n.check()

	return n.application.response(cacheType)
}

func (n *Nginx) check() bool {
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
