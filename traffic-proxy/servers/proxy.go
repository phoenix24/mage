package servers

import (
	"fmt"
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


func connProxy(in <-chan *net.TCPConn, out chan<- *net.TCPConn) {
	for conn := range in {
		connHandler()
	}
}

func connHandler(connPair *ConnPair) {
	var pConn, bConn = connPair.pConn, connPair.bConn
	defer pConn.Close()
	defer bConn.Close()

	log.Printf("serving %s -> %s\n", pConn.RemoteAddr().String(), bConn.RemoteAddr().String())
	for {
		var pbuf = make([]byte, 4096)
		var pread, _ = pConn.Read(pbuf)
		fmt.Println("1. read/ wrote: ", pread, len(pbuf))

		var pwritten, _ = bConn.Write(pbuf[:pread])
		fmt.Println("1. write complete: ", pwritten)

		var bbuf = make([]byte, 4096)
		var bread, _ = bConn.Read(bbuf)
		fmt.Println("2. read/ wrote: ", bread, len(bbuf))

		var bwritten, _ = pConn.Write(bbuf[:bread])
		fmt.Println("2. write complete: ", bwritten)
	}
}

func connCleaner(connPair *ConnPair) {

}


func (s *ProxyServer) ListenAndServe() error {
	var message = fmt.Sprintf("\nstarting server : %s\n", s.name) +
		fmt.Sprintf("    with mode   : %s\n", s.mode) +
		fmt.Sprintf("    with route  : %s -> %s\n", s.source, s.backend) +
		fmt.Sprintf("    with stores : %s\n", s.storage)
	log.Println(message)

	//todo: mode -> tcp-proxy (source <-> backend)
	// connection manager.
	// connection manager.

	//todo: mode -> tcp-proxy + record traffic (in-memory).
	//todo: storage -> in-memory
	//todo: storage -> pubsub queue

	//todo: mode -> tcp-proxy - replay traffic only.
	//todo: request matcher (bytestream)?
	//todo: request matcher (parsed request)?

	var port = ":" + s.source.Port()
	var listen, err = net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("error listening on the port: ", port)
	}
	defer listen.Close()

	var pendingConns = make(chan *net.TCPConn)
	var completeConns = make(chan *net.TCPConn)

	go connProxy(pendingConns, completeConns)
	go connCleaner(completeConns)

	for {
		var proxyConn, _ = listen.Accept()
		var backendTCP, _ = net.ResolveTCPAddr("tcp", s.backend.HostPort())
		var backendConn, _ = net.DialTCP("tcp", nil, backendTCP)
		go connHandler(&ConnPair{pConn: proxyConn, bConn: backendConn})
	}
	return nil
}
