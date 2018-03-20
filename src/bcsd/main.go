package main

import (
	"BCSD/Server"
)

func main() {
	srv := new(server.Service)
	srv.Start()
}