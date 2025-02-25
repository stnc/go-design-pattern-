package main

import (
	"fmt"
	pac "proxy/app"
	c "proxy/app/cacheRepo"
)

func main() {

	repoProxyServer := pac.RepoProxy()

	fmt.Println(repoProxyServer.RepoRequest())
	c.TesGetByUsername()
}
