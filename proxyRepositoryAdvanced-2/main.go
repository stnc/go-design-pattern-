package main

import (
	"fmt"
	pac "proxy/app"
)

func main() {

	repoProxyServer := pac.RepoProxy()

	fmt.Println(repoProxyServer.RepoRequest())

}
