package main

import (
	server "github.com/kzemlyak/microchelik/user/server"
)

func main() {
	server.InitHttpServer()
	server.InitGrpcServer()
}
