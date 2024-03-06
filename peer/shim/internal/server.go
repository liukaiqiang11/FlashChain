package internal

import (
	"crypto/tls"
	"errors"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

const (
	serverMinInterval = time.Duration(1) * time.Minute
	connectionTimeout = 5 * time.Second
)

type Server struct {
	Listener net.Listener
	Server   *grpc.Server
}

func (s *Server) Start() error {
	if s.Listener == nil {
		return errors.New("nil listener")
	}

	if s.Server == nil {
		return errors.New("nil server")
	}

	return s.Server.Serve(s.Listener)
}

func (s *Server) Stop() {
	if s.Server != nil {
		s.Server.Stop()
	}
}

func NewServer(
	address string,
	tlsConf *tls.Config,
) (*Server, error) {
	if address == "" {
		return nil, errors.New("server listen address not provided")
	}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	var serverOpts []grpc.ServerOption

	serverKeepAliveParameters := keepalive.ServerParameters{
		Time:    1 * time.Minute,
		Timeout: 20 * time.Second,
	}
	serverOpts = append(serverOpts, grpc.KeepaliveParams(serverKeepAliveParameters))

	if tlsConf != nil {
		creds := credentials.NewTLS(tlsConf)
		serverOpts = append(serverOpts, grpc.Creds(creds))
	}

	serverOpts = append(serverOpts, grpc.MaxSendMsgSize(maxSendMessageSize))
	serverOpts = append(serverOpts, grpc.MaxRecvMsgSize(maxRecvMessageSize))

	kep := keepalive.EnforcementPolicy{
		MinTime:             serverMinInterval,
		PermitWithoutStream: true,
	}
	serverOpts = append(serverOpts, grpc.KeepaliveEnforcementPolicy(kep))

	serverOpts = append(serverOpts, grpc.ConnectionTimeout(connectionTimeout))

	server := grpc.NewServer(serverOpts...)

	return &Server{Listener: listener, Server: server}, nil
}
