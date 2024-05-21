package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

var POISON = struct{}{}

type Server struct {
	listener net.Listener
	exitC    chan struct{}
}

func NewServer() Server {
	return Server{
		exitC: make(chan struct{}, 1),
	}
}

func (p *Server) Bind(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	p.listener = l
	return nil
}

func (p *Server) Start() error {
	if p.listener == nil {
		return fmt.Errorf("Server is not binding")
	}
	defer p.listener.Close()

LOOP:
	for {
		select {
		case <-p.exitC:
			break LOOP
		default:
			p.accept()
		}
	}

	return nil
}

func (p *Server) Stop() {
	p.exitC <- POISON
}

func (p *Server) accept() {
	conn, err := p.listener.Accept()
	if err != nil {
		fmt.Printf("Error in accept: %v\n", err)
		return
	}

	fmt.Println("New client connect.")
	go p.handle(conn)
}

func (p *Server) handle(c net.Conn) {
	defer c.Close()

	var (
		reader = bufio.NewReader(c)
		bytes  = make([]byte, 1024)
	)

	for {
		_, err := reader.Read(bytes)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Client close.\n")
			} else {
				fmt.Printf("Error in handle: %v\n", err)
			}
			break
		}

		fmt.Printf("Read from client: %v\n", string(bytes))
		c.Write([]byte("Yes I see"))
	}
}
