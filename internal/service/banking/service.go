package banking

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"komek/internal/service/banking/pb"
	"time"
)

type Banking struct {
	client pb.BankingClient
}

func New(addr string, enableTLS bool) (*Banking, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	creds := insecure.NewCredentials()
	if enableTLS {
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			return nil, fmt.Errorf("enable tls -> system roots: %w", err)
		}
		creds = credentials.NewTLS(&tls.Config{RootCAs: systemRoots})
	}
	client, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, fmt.Errorf("grpc -> dial: %w", err)
	}

	return &Banking{pb.NewBankingClient(client)}, nil
}