package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Starting server...")

	server := NewServer()
	if err := server.Bind(":8081"); err != nil {
		fmt.Println(err)
		return
	}

	go server.Start()
	fmt.Println("Server started!")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-sig

	server.Stop()

	fmt.Println("Wait for server stop.")
	time.Sleep(3 * time.Second)
}
