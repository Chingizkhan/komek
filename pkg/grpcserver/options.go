package grpcserver

type Option func(*Server)

func W(smth string) Option {
	return func(s *Server) {
		//s.smth = smth
	}
}
