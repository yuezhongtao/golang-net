package main

import (
	"github.com/yuezhongtao/golang-net/tcp/client"
	"github.com/yuezhongtao/golang-net/tcp/server"
)

func main() {
	go server.StartServer()
	client.StartClient()
}
