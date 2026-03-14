package server

import (
	"bufio"
	"io"
	"net"
)

type Server struct {
	addr string
}

func New(addr string) *Server {
	return &Server{addr: addr}
}

// Start はTCPサーバーを起動し、接続受け付けループを開始する
func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go s.handleConn(conn)
	}
}

// handleConn は1接続ごとの読み書きを処理する
func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			return
		}

		_, _ = conn.Write([]byte("received: " + line))
	}
}
