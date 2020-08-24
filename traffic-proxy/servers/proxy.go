package servers

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
	"traffic-proxy/configs"
)

type SrcID string

const (
	SrcClient = "client"
	SrcServer = "server"
)

type MessageProtocol uint8

const (
	HTTP = 0
	MYSQL
	PGSQL
	REDIS
)

type Message struct {
	data  []byte
	time  time.Time
	srcIp net.IP
	srcId SrcID
}

type MessagePair struct {
	ID       string //todo: must be a uuid (type)?
	Request  Message
	Response Message
	Protocol MessageProtocol
}

type ConnPair struct {
	pConn    net.Conn
	bConn    net.Conn
	messages []Message
}

type ProxyServer struct {
	name    string
	mode    configs.Mode
	sinks   []io.Writer
	source  configs.Address
	backend configs.Address
}

func (s *ProxyServer) broker(dst, src net.Conn, srcChan chan<- struct{}, meta string) {
	var sinks = append(s.sinks, dst)
	_, err := io.Copy(io.MultiWriter(sinks...), src)
	if err != nil {
		log.Printf("%s, copy error: %s", meta, err)
	}
	if err := src.Close(); err != nil {
		log.Printf("%s, close error: %s", meta, err)
	}
	srcChan <- struct{}{}
}

func (s *ProxyServer) handler(clientConn, serverConn *net.TCPConn) {
	var serverClosed = make(chan struct{}, 1)
	var clientClosed = make(chan struct{}, 1)

	go s.broker(serverConn, clientConn, clientClosed, "client")
	go s.broker(clientConn, serverConn, serverClosed, "server")

	var waitFor chan struct{}
	select {
	case <-clientClosed:
		serverConn.SetLinger(0)
		serverConn.Close()
		waitFor = serverClosed
	case <-serverClosed:
		clientConn.Close()
		waitFor = clientClosed
	}

	<-waitFor
}

func (s *ProxyServer) ListenAndServe() error {
	var message = fmt.Sprintf("\nstarting server : %s\n", s.name) +
		fmt.Sprintf("    with mode   : %s\n", s.mode) +
		fmt.Sprintf("    with route  : %s -> %s\n", s.source, s.backend) +
		fmt.Sprintf("    with sinks  : %s", s.sinks)
	log.Println(message)

	//todo: make file to support cross-os builds.

	//todo: mode -> proxy (source <> backend)
	// sinks - null, console etc.

	//todo: mode -> record traffic
	// sinks - file, redis, inmemory, kafka, pulsar, database etc.

	//todo: mode -> replay traffic
	// request matcher (bytestream)?
	// request matcher (parsed request)?

	var listen, err = net.Listen("tcp", s.source.HostPort())
	if err != nil {
		log.Fatalln("error listening on the port: ", s.source.HostPort())
	}
	defer listen.Close()

	for {
		var clientConn, _ = listen.Accept()
		var serverConn, err = net.Dial("tcp", s.backend.HostPort())
		if err != nil {
			log.Println("failed to connect to backend server: ", err)
			clientConn.Close()
		} else {
			log.Printf("serving %s -> %s\n", clientConn.RemoteAddr().String(), serverConn.RemoteAddr().String())
			go s.handler(clientConn.(*net.TCPConn), serverConn.(*net.TCPConn))
		}
	}
	return nil
}
