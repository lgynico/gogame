package main

import (
	"fmt"
	"net"
)

type Client struct {
	conn net.Conn
}

func NewClient() Client {
	return Client{}
}

func (p *Client) Connect(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	p.conn = conn
	return nil
}

func (p *Client) Start() error {
	defer p.conn.Close()

	p.Write("Hello Server.")
	p.Write("I'm a new client.")
	p.Write("Goodbye!")

	fmt.Printf("Read from server: %s\n", p.Read())
	return nil
}

func (p *Client) Write(text string) error {
	_, err := p.conn.Write([]byte(text))
	return err
}

func (p *Client) Read() string {
	bytes := make([]byte, 1024)
	if _, err := p.conn.Read(bytes); err != nil {
		return ""
	}

	return string(bytes)
}
