package servers

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"traffic-proxy/configs"
)

type ProxyServer struct {
}

func (s *ProxyServer) handler(conn net.Conn) {
	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())
	for {
		var data, _ = bufio.NewReader(conn).ReadString('\n')
		conn.Write([]byte("hello world! " + data))
	}
}

func (s *ProxyServer) ListenAndServe(config configs.ServerConfig) error {
	var message = fmt.Sprintf("\nstarting server : %s\n", config.Name) +
		fmt.Sprintf("    with mode   : %s\n", config.Mode) +
		fmt.Sprintf("    with route  : %s -> %s\n", config.Source, config.Backend) +
		fmt.Sprintf("    with stores : %s\n", config.Storage)
	log.Println(message)

	//mode
	//source
	//backend
	//storage

	var listen, _ = net.Listen("tcp4", config.Source.Port())
	defer listen.Close()

	go func () {
		var conn, _ = listen.Accept()
		go s.handler(conn)
	}()
	return nil
}
