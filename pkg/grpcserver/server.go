package grpcserver

import (
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	addr   string
	notify chan error
	server *grpc.Server
}

func New(grpcServer *grpc.Server, port string, opts ...Option) *Server {
	addr := ":" + port

	s := &Server{
		server: grpcServer,
		addr:   addr,
		notify: make(chan error, 1),
	}

	for _, opt := range opts {
		opt(s)
	}

	s.Start()

	return s
}

func (s *Server) Start() {
	go func() {
		listener, err := net.Listen("tcp", s.addr)
		if err != nil {
			s.notify <- err
			close(s.notify)
			return
		}

		s.notify <- s.server.Serve(listener)
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() {
	s.server.GracefulStop()
}
