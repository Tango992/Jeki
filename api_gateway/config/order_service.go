package config

import (
	pb "api-gateway/pb/orderpb"
	"crypto/tls"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func InitOrderServiceClient() (*grpc.ClientConn, pb.OrderServiceClient) {
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(os.Getenv("ORDER_SERVICE_URI"), opts...)
	if err != nil {
		log.Fatal(err)
	}
	return conn, pb.NewOrderServiceClient(conn)
}