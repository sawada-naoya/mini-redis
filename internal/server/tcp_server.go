package server

import (
	"bufio"
	"io"
	"log"
	"net"

	"github.com/sawada-naoya/mini-redis/internal/handler"
	"github.com/sawada-naoya/mini-redis/internal/protocol"
	"github.com/sawada-naoya/mini-redis/internal/store"
)

type Server struct {
	addr    string
	handler *handler.Handler
}

func New(addr string) *Server {
	st := store.New()
	h := handler.New(st)

	return &Server{
		addr:    addr,
		handler: h,
	}
}

// Start はTCPサーバーを起動し、接続受け付けループを開始する
func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	log.Printf("server started on %s", s.addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		log.Printf("accepted connection from %s", conn.RemoteAddr())
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

		res := s.handler.Execute(cmd)

		_, writeErr := conn.Write([]byte(res))
		if writeErr != nil {
			log.Printf("write error to %s: %v", conn.RemoteAddr(), writeErr)
			return
		}
	}
}
