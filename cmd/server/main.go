package main

import "github.com/metajar/gofor/internal/gofor"

func main() {
	hosts := []string{
		"https://jsonplaceholder.typicode.com",
		"https://jsonplaceholder.typicode.com",
	}
	g := gofor.New(hosts, "8888")
	g.StartServer()

}
