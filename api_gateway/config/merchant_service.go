package config

import (
	pb "api-gateway/pb/merchantpb"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitMerchantServiceClient() (*grpc.ClientConn, pb.MerchantClient) {
	// creds := credentials.NewTLS(&tls.Config{
	// 	InsecureSkipVerify: true,
	// })

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(os.Getenv("MERCHANT_SERVICE_URI"), opts...)
	if err != nil {
		log.Fatal(err)
	}
	return conn, pb.NewMerchantClient(conn)
}