package main

import (
	server "github.com/kzemlyak/microchelik/chat/server"
)

func main() {
	server.InitHttpServer()
	server.InitGrpcServer()
}
