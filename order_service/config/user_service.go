package config

import (
	"crypto/tls"
	"log"
	"order-service/pb/userpb"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func InitUserServiceClient() (*grpc.ClientConn, userpb.UserClient) {
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(os.Getenv("USER_SERVICE_URI"), opts...)
	if err != nil {
		log.Fatal(err)
	}
	return conn, userpb.NewUserClient(conn)
}