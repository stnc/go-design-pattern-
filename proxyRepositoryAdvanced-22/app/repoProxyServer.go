package dll

import "fmt"

type Repo struct {
	//	application *Application
	manuelClose bool //only database
	//	rateLimiter map[string]int
}

type Application struct {
}

func RepoProxy() *Repo {
	return &Repo{
		//	application: &Application{},
		manuelClose: false, //only database
		//	rateLimiter: make(map[string]int),
	}
}

func (n *Repo) RepoRequest() Repositories {
	cacheType := n.check()
	if cacheType == true {
		fmt.Println("cacheRepo1")
		return cacheRepo()
	} else {
		fmt.Println("postRepo1")
		return postRepo()
	}
}

func (n *Repo) check() bool {
	cacheOK := true
	////override cache
	if n.manuelClose == true {
		fmt.Println("close")
		return false //cache close
	} else {
		fmt.Println("open")
		if cacheOK == true {
			fmt.Println("open1")
			return true
		} else {
			fmt.Println("open2")
			return false
		}
	}
}
