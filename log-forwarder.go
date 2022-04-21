package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

const (
	MAX_TRIES = 60
)

func main() {
	log.SetOutput(os.Stderr)
	log.SetFlags(0)
	log.SetPrefix("")

	if len(os.Args) < 4 {
		log.Fatalf("Usage: %s <fifo> <host> <port>\n", os.Args[0])
	}
	fifo := os.Args[1]

	addr := net.JoinHostPort(os.Args[2], os.Args[3])

	var conn *net.TCPConn
	tries := 0
	for conn == nil {
		if tries == MAX_TRIES {
			log.Fatalf("Not connected after %s attempts", MAX_TRIES)
		}
		tries += 1

		c, err := net.Dial("tcp", addr)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		conn = c.(*net.TCPConn)
	}
	defer conn.Close()

	if err := conn.CloseRead(); err != nil {
		log.Fatal(err)
	}

	if err := conn.SetKeepAlive(true); err != nil {
		log.Fatal(err)
	}

	p, _ := time.ParseDuration("30s")
	if err := conn.SetKeepAlivePeriod(p); err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected to %s\n", addr)

	f, err := os.Open(fifo)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.Printf("Opened %s for reading\n", fifo)

	if _, err = io.Copy(conn, f); err != nil {
		log.Fatal(err)
	}
}
