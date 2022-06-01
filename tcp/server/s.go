package server

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var port = flag.Int("port", 18080, "the server port")

func StartServer() {
	flag.Parse()
	log.Printf("start listening...")

	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf(err.Error())
		}
		log.Printf("connection from: %s", conn.RemoteAddr())
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("connection maybe closed, error: %s", err)
		}
	}(conn)

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatalf("read from connection error: %s", err)
		}
		log.Printf("receive from client: %s", string(buf[:n]))

		_, err = conn.Write([]byte(fmt.Sprintf("server has received: %s", string(buf[:n]))))
		if err != nil {
			log.Fatalf("send to client error: %s", err)
		}

	}

}
