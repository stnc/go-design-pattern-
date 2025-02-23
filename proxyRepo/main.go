package main

import (
	"fmt"
	pac "proxy/app"
)

func main() {

	nginxServer := pac.NginxServer()

	fmt.Println(nginxServer.RepoRequest())

}
