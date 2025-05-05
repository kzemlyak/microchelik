package main

import (
	server "github.com/kzemlyak/microchelik/social/server"
)

func main() {
	server.InitHttpServer()
	server.InitGrpcServer()
}
