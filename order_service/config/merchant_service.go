package config

import (
	"log"
	"order-service/pb/merchantpb"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitMerchantServiceClient() (*grpc.ClientConn, merchantpb.MerchantClient) {
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
	return conn, merchantpb.NewMerchantClient(conn)
}