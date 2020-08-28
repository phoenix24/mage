package servers

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
	"traffic-proxy/common"
	"traffic-proxy/configs"
)

type SrcID string

const (
	SrcClient = "client"
	SrcServer = "server"
)

type ConnHandler struct {
	sink     chan *common.Packet
	client   *net.TCPConn
	remote   *net.TCPConn
	protocol *string
}

type ProxyServer struct {
	name     string
	sink     chan *common.Packet
	mode     configs.Mode
	source   configs.Address
	remote   configs.Address
	protocol string
}

func (c *ConnHandler) Write(payload []byte) (n int, err error) {
	c.sink <- &common.Packet{Data: payload, Time: time.Now(), Protocol: c.protocol}
	return len(payload), nil
}

func (c *ConnHandler) broker(dst, src net.Conn, srcChan chan<- struct{}, meta string) {
	_, err := io.Copy(io.MultiWriter(c, dst), src)
	if err != nil {
		log.Printf("%s, copy error: %s", meta, err)
	}
	if err := src.Close(); err != nil {
		log.Printf("%s, close error: %s", meta, err)
	}
	srcChan <- struct{}{}
}

func (c *ConnHandler) handle() {
	var remoteClosed = make(chan struct{}, 1)
	var clientClosed = make(chan struct{}, 1)

	go c.broker(c.client, c.remote, remoteClosed, "remote")
	go c.broker(c.remote, c.client, clientClosed, "client")

	select {
	case <- clientClosed:
		c.remote.SetLinger(0)
		c.remote.Close()

	case <- remoteClosed:
		c.client.Close()
	}
}

func (p *ProxyServer) ListenAndServe() error {
	var message = fmt.Sprintf("\nstarting server : %s\n", p.name) +
		fmt.Sprintf("    with mode   : %s\n", p.mode) +
		fmt.Sprintf("    with route  : %s -> %s\n", p.source, p.remote)
	log.Println(message)

	var listen, err = net.Listen("tcp", p.source.HostPort())
	if err != nil {
		log.Fatalln("error listening on the port: ", p.source.HostPort())
	}
	defer listen.Close()

	for {
		var client, _ = listen.Accept()
		var remote, err = net.Dial("tcp", p.remote.HostPort())
		if err != nil {
			log.Println("failed to connect to remote server: ", err)
			client.Close()
		} else {

			log.Printf("serving %s -> %s\n", client.RemoteAddr().String(), remote.RemoteAddr().String())
			var handler = &ConnHandler{
				sink: p.sink,
				client: client.(*net.TCPConn),
				remote: remote.(*net.TCPConn),
				protocol: &p.protocol,
			}
			go handler.handle()
		}
	}
	return nil
}
