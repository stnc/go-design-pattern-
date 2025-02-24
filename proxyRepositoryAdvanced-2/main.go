package main

import (
	"fmt"
	pac "proxy/app"
)

func main() {

	repoProxyServer := pac.RepoProxyServer()

	fmt.Println(repoProxyServer.RepoRequest())

}
