package config

import (
	"api-gateway/pb"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitGrpc() (*grpc.ClientConn, pb.UserClient) {
	// creds := credentials.NewTLS(&tls.Config{
	// 	InsecureSkipVerify: true,
	// })

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(os.Getenv("GRPC_URI"), opts...)
	if err != nil {
		log.Fatal(err)
	}
	return conn, pb.NewUserClient(conn)
}