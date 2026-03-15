package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/sawada-naoya/mini-redis/internal/protocol"
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
	defer func() {
		log.Printf("closed connection from %s", conn.RemoteAddr())
		conn.Close()
	}()

	reader := bufio.NewReader(conn)

	for {
		// クライアントから1行受け取る
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Printf("read error from %s: %v", conn.RemoteAddr(), err)
			return
		}

		cmd, err := protocol.ParseLine(line)
		if err != nil {
			_, _ = conn.Write([]byte("ERR invalid command\n"))
			continue
		}

		res := fmt.Sprintf("command=%s args=%v\n", cmd.Name, cmd.Args)
		_, writeErr := conn.Write([]byte(res))
		if writeErr != nil {
			log.Printf("write error to %s: %v", conn.RemoteAddr(), writeErr)
			return
		}
	}
}
