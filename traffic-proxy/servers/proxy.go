package servers

import (
	"fmt"
	"io"
	"log"
	"net"
	"traffic-proxy/configs"
)

type ConnPair struct {
	pConn net.Conn
	bConn *net.TCPConn
}

type ProxyServer struct {
	name    string
	mode    configs.Mode
	source  configs.Address
	backend configs.Address
	storage configs.StorageURL
}

func broker(dst, src net.Conn, srcClosed chan struct{}) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Printf("Copy error: %s", err)
	}
	if err := src.Close(); err != nil {
		log.Printf("Close error: %s", err)
	}
	srcClosed <- struct{}{}
}

func (s *ProxyServer) handler(clientConn, serverConn net.Conn) {
	var serverClosed = make(chan struct{}, 1)
	var clientClosed = make(chan struct{}, 1)

	go broker(serverConn, clientConn, clientClosed)
	go broker(clientConn, serverConn, serverClosed)

	var waitFor chan struct{}
	select {
	case <- clientClosed:
		serverConn.Close()
		waitFor = serverClosed
	case <- serverClosed:
		clientConn.Close()
		waitFor = clientClosed
	}

	<-waitFor
}

func (s *ProxyServer) ListenAndServe() error {
	var message =
		fmt.Sprintf("\nstarting server : %s\n", s.name) +
		fmt.Sprintf("    with mode   : %s\n", s.mode) +
		fmt.Sprintf("    with route  : %s -> %s\n", s.source, s.backend) +
		fmt.Sprintf("    with stores : %s", s.storage)
	log.Println(message)

	//todo:
	// make file to support cross-os builds.

	//todo: mode -> tcp-proxy (source <-> backend)
	// connection manager.
	// sinks - null,file,console,kafka,pulsar,etc.

	//todo: mode -> tcp-proxy + record traffic (in-memory).
	//todo: storage -> in-memory
	//todo: storage -> pubsub queue

	//todo: mode -> tcp-proxy - replay traffic only.
	//todo: request matcher (bytestream)?
	//todo: request matcher (parsed request)?

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
		}

		log.Printf("serving %s -> %s\n", clientConn.RemoteAddr().String(), serverConn.RemoteAddr().String())
		go s.handler(clientConn, serverConn)
	}
	return nil
}
