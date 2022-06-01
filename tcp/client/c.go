package client

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func StartClient() {
	log.Printf("start client ...")
	conn, err := net.Dial("tcp", "localhost:18080")
	if err != nil {
		log.Fatalf("connect to server failed: %s\n", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {

		str, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("read from stdin err=%v", err)
			return
		}
		//输入exit退出。
		if strings.Trim(str, "\r\n") == "exit" {
			log.Printf("client exit")
			return
		}

		// 把数据发送给服务器
		n, err := conn.Write([]byte(str))
		if err != nil {
			log.Printf("send data error err=%v", err)
			return
		}
		log.Printf("send %d bit data。\n", n)

		buf := make([]byte, 1024)
		n2, err := conn.Read(buf)
		if err != nil {
			log.Fatalf("read from connection error: %s", err)
		}
		log.Printf("receive from server: %s", string(buf[:n2]))

	}

}
